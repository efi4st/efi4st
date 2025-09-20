/**
 * Author:    Admiral Helmut
 * Created:   12.06.2019
 *
 * (C)
 **/

package routes

import (
	"fmt"
	"github.com/efi4st/efi4st/classes"
	"github.com/efi4st/efi4st/dbprovider"
	_ "github.com/go-sql-driver/mysql"
	"github.com/kataras/iris/v12"
	"io"
	"log"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"strings"
	"time"
)

/////////////
//
// ChecklistTemplates
//
/////////////


var (
	// wohin du die hochgeladenen HTML/MD-Dateien schreibst
	docAssetBase = "/var/app/docs/checklist-assets"
	// (nur falls du es woanders brauchst)
	StaticFSDir  = "/var/app/static"
)

type DocAssets struct {
	CoverPath  string // leer, wenn nicht vorhanden
	FooterPath string // leer, wenn nicht vorhanden
}


func SMSChecklistTemplates(ctx iris.Context) {
	templates := dbprovider.GetDBManager().GetAllChecklistTemplates()
	ctx.ViewData("templateList", templates)
	ctx.View("sms_checklistTemplates.html")
}

func CreateSMSChecklistTemplate(ctx iris.Context) {
	ctx.View("sms_createChecklistTemplate.html")
}

func AddSMSChecklistTemplate(ctx iris.Context) {
	name := ctx.PostValue("Name")
	description := ctx.PostValue("Description")

	err := dbprovider.GetDBManager().AddChecklistTemplate(&classes.Sms_ChecklistTemplate{
		Name: name,
		Description: description,
	})
	if err != nil {
		ctx.ViewData("error", "Error: Could not create checklist template")
	}
	templates := dbprovider.GetDBManager().GetAllChecklistTemplates()
	ctx.ViewData("templateList", templates)
	ctx.View("sms_checklistTemplates.html")
}

func ShowSMSChecklistTemplate(ctx iris.Context) {
	idStr := ctx.Params().Get("id")
	id, _ := strconv.Atoi(idStr)

	template := dbprovider.GetDBManager().GetChecklistTemplateByID(id)
	items := dbprovider.GetDBManager().GetChecklistTemplateItems(id)
	artefactTypes := dbprovider.GetDBManager().GetSMSArtefactTypes()

	checkDefs, err := dbprovider.GetDBManager().GetAllDeviceCheckDefinitions()
	if err != nil {
		log.Println("Fehler beim Abrufen der Check Definitions:", err)
	}

	scheme := "http"
	if ctx.Request().TLS != nil { scheme = "https" }
	assetInfo := GetDocAssetInfoForTemplate(id, scheme, ctx.Host())

	ctx.ViewData("template", template)
	ctx.ViewData("items", items)
	ctx.ViewData("artefactTypes", artefactTypes)
	ctx.ViewData("checkDefinitions", checkDefs)
	ctx.ViewData("docAssets", assetInfo) // ← pass to view

	ctx.View("sms_showChecklistTemplate.html")
}


func RemoveSMSChecklistTemplate(ctx iris.Context) {
	idStr := ctx.Params().Get("id")
	id, _ := strconv.Atoi(idStr)

	_ = dbprovider.GetDBManager().DeleteChecklistTemplateByID(id)

	templates := dbprovider.GetDBManager().GetAllChecklistTemplates()
	ctx.ViewData("templateList", templates)
	ctx.View("sms_checklistTemplates.html")
}

/////////////
//
// ChecklistTemplateItem
//
/////////////
func AddSMSChecklistTemplateItem(ctx iris.Context) {
	templateID, _ := strconv.Atoi(ctx.PostValue("ChecklistTemplateID"))
	checkDefinitionID := ctx.PostValueIntDefault("CheckDefinitionID", 0)
	artefactTypeID := ctx.PostValueIntDefault("ArtefactTypeID", 0)
	scope := ctx.PostValue("TargetScope")
	expected := ctx.PostValue("ExpectedValue")
	optional := ctx.PostValue("Optional") == "on"

	var checkIDPtr *int = nil
	if checkDefinitionID != 0 {
		checkIDPtr = &checkDefinitionID
	}
	var artefactIDPtr *int = nil
	if artefactTypeID != 0 {
		artefactIDPtr = &artefactTypeID
	}

	item := &classes.Sms_ChecklistTemplateItem{
		ChecklistTemplateID: templateID,
		CheckDefinitionID:   checkIDPtr,
		ArtefactTypeID:      artefactIDPtr,
		TargetScope:         scope,
		ExpectedValue:       expected,
		Optional:            optional,
	}

	err := dbprovider.GetDBManager().AddChecklistTemplateItem(item)
	if err != nil {
		log.Println("Fehler beim Hinzufügen des Checklist-Items:", err)
		ctx.ViewData("error", "Fehler beim Hinzufügen des Checklist-Items: "+err.Error())

		// Wiederlade die Seite mit Template + Items zur Anzeige des Fehlers
		template := dbprovider.GetDBManager().GetChecklistTemplateByID(templateID)
		items := dbprovider.GetDBManager().GetChecklistTemplateItems(templateID)

		ctx.ViewData("template", template)
		ctx.ViewData("items", items)
		ctx.View("sms_showChecklistTemplate.html")
		return
	}

	// Wenn kein Fehler: Redirect zur Übersicht
	http.Redirect(ctx.ResponseWriter(), ctx.Request(), fmt.Sprintf("/sms_checklistTemplate/show/%d", templateID), http.StatusFound)
}

func RemoveSMSChecklistTemplateItem(ctx iris.Context) {
	id, _ := strconv.Atoi(ctx.Params().Get("id"))
	templateID, _ := strconv.Atoi(ctx.URLParam("template_id"))
	_ = dbprovider.GetDBManager().DeleteChecklistTemplateItemByID(id)
	ctx.Redirect(fmt.Sprintf("/sms_checklistTemplate/show/%d", templateID))
}

///////////////
//
//// ChecklistInstance
//
//////////////

func GenerateChecklistInstanceForProject(ctx iris.Context) {
	projectID, _ := strconv.Atoi(ctx.Params().Get("project_id"))
	templateID, _ := strconv.Atoi(ctx.PostValue("ChecklistTemplateID"))

	instance := &classes.Sms_ChecklistInstance{
		ProjectID:           &projectID,
		DeviceID:            nil,
		ChecklistTemplateID: templateID,
		GeneratedBy:         "system", // ← Beispielwert
		Status:              "open",
	}

	err := dbprovider.GetDBManager().AddChecklistInstance(instance)
	if err != nil {
		log.Printf("❌ Fehler beim Instanzieren der Checklist (Template %d, Project %d): %v", templateID, projectID, err)
		ctx.ViewData("error", fmt.Sprintf("Fehler beim Erzeugen der Checkliste: %v", err))

		// Lade Projekt erneut, damit wir zur Ursprungsseite zurückkehren können
		project := dbprovider.GetDBManager().GetSMSProjectInfo(projectID)
		ctx.ViewData("project", project)

		// Optional: zeige Fehlermeldung direkt in Projekt-Ansicht oder eigene Error-View
		ctx.View("sms_showProject.html")
		return
	}

	ctx.Redirect(fmt.Sprintf("/sms_projects/show/%d", projectID))
}

func GenerateChecklistInstanceForDevice(ctx iris.Context) {
	deviceID, _ := strconv.Atoi(ctx.Params().Get("device_id"))
	templateID, _ := strconv.Atoi(ctx.PostValue("ChecklistTemplateID"))

	err := dbprovider.GetDBManager().AddChecklistInstance(&classes.Sms_ChecklistInstance{
		ProjectID:           nil,
		DeviceID:            &deviceID,
		ChecklistTemplateID: templateID,
		Status:              "open",
	})
	if err != nil {
		ctx.ViewData("error", "Failed to create checklist instance")
	}
	ctx.Redirect(fmt.Sprintf("/sms_devices/show/%d", deviceID))
}

func ShowChecklistInstance(ctx iris.Context) {
	id, _ := strconv.Atoi(ctx.Params().Get("id"))

	inst := dbprovider.GetDBManager().GetChecklistInstanceByID(id)
	if inst == nil {
		ctx.ViewData("error", "Checklist instance not found")
		ctx.View("sms_checklistTemplates.html")
		return
	}

	items := dbprovider.GetDBManager().GetChecklistItemInstancesWithDefinition(id)

	// Projekt-Kontext: passende DeviceInstances listen
	if inst.ProjectID != nil {
		for i := range items {
			if items[i].CheckDefinitionID != nil && items[i].DeviceTypeID != nil {
				candidates := dbprovider.GetDBManager().
					GetDeviceInstancesForProjectAndDeviceType(*inst.ProjectID, *items[i].DeviceTypeID)

				var matches []classes.MatchingDevice
				for _, c := range candidates {
					if matchesApplicable(c.DeviceVersion, items[i].ApplicableVersions) {
						matches = append(matches, c)
					}
				}
				items[i].MatchingDevices = matches
			}
			// ▼ Für Projekt-Kontext keine Device-Anwendbarkeit: explizit "none"
			items[i].AppliesToThisDevice = nil
			items[i].AppliesToThisDeviceStr = "none"
		}
	}

	// Device-Kontext: diesen Device-Typ + Version prüfen
	if inst.DeviceID != nil {
		dbasic, err := dbprovider.GetDBManager().GetDeviceBasicByID(*inst.DeviceID)
		if err == nil && dbasic != nil {
			for i := range items {
				items[i].DeviceContextTypeName = dbasic.DeviceType
				items[i].DeviceContextVersion  = dbasic.Version

				if items[i].CheckDefinitionID != nil {
					var typeOK bool
					if items[i].DeviceTypeID != nil {
						typeOK = (*items[i].DeviceTypeID == dbasic.DeviceTypeID)
					} else if items[i].DeviceTypeName != "" {
						typeOK = strings.EqualFold(items[i].DeviceTypeName, dbasic.DeviceType)
					} else {
						typeOK = false
					}
					versOK := matchesApplicable(dbasic.Version, items[i].ApplicableVersions)

					res := typeOK && versOK
					items[i].AppliesToThisDevice = &res
					// ▼ NEU: String-Repräsentation setzen
					if res {
						items[i].AppliesToThisDeviceStr = "true"
					} else {
						items[i].AppliesToThisDeviceStr = "false"
					}

					log.Printf("Check apply? inst=%d item=%d typeOK=%v (need %v/%s have %d/%s) versOK=%v (need '%s' have '%s')",
						inst.ChecklistInstanceID, items[i].ChecklistItemInstanceID,
						typeOK,
						valueOr(items[i].DeviceTypeID, 0), items[i].DeviceTypeName,
						dbasic.DeviceTypeID, dbasic.DeviceType,
						versOK,
						items[i].ApplicableVersions, dbasic.Version,
					)
				} else {
					// ▼ NEU: explizit „none“ setzen, wenn keine Definition
					items[i].AppliesToThisDevice = nil
					items[i].AppliesToThisDeviceStr = "none"
				}
			}
		}
	}
	mgr := dbprovider.GetDBManager()

	mgr.PostRenderChecklistItems(inst, items)

	ctx.ViewData("instance", inst)
	ctx.ViewData("items", items)
	ctx.View("sms_showChecklistInstance.html")
}

// kleiner helper:
func valueOr(p *int, def int) int {
	if p == nil { return def }
	return *p
}

// 'all' oder CSV-Liste exakter Versionsstrings (einfacher Ansatz)
func matchesApplicable(version, applicable string) bool {
	if applicable == "" || strings.EqualFold(applicable, "all") {
		return true
	}
	for _, tok := range strings.Split(applicable, ",") {
		if strings.EqualFold(strings.TrimSpace(tok), strings.TrimSpace(version)) {
			return true
		}
	}
	return false
}

func DeleteChecklistInstance(ctx iris.Context) {
	id, _ := strconv.Atoi(ctx.Params().Get("id"))
	dbprovider.GetDBManager().DeleteChecklistInstanceByID(id)
	ctx.Redirect("/sms_checklistTemplates")
}

func MarkChecklistInstanceStatus(ctx iris.Context) {
	id, _ := strconv.Atoi(ctx.Params().Get("id"))
	status := ctx.URLParam("status")
	_ = dbprovider.GetDBManager().UpdateChecklistInstanceStatus(id, status)
	ctx.Redirect(fmt.Sprintf("/sms_checklistInstance/show/%d", id))
}


func UpdateChecklistItemInstance(ctx iris.Context) {
	itemID, _ := strconv.Atoi(ctx.PostValue("ChecklistItemInstanceID"))
	checklistInstanceID, _ := strconv.Atoi(ctx.PostValue("ChecklistInstanceID"))
	isOKStr := ctx.PostValue("IsOK")
	actualValue := ctx.PostValue("ActualValue")
	comment := ctx.PostValue("Comment")

	var isOK *bool = nil
	if isOKStr != "" {
		val := isOKStr == "true" || isOKStr == "on"
		isOK = &val
	}

	item := &classes.Sms_ChecklistItemInstance{
		ChecklistItemInstanceID: itemID,
		IsOK:                    isOK,
		ActualValue:             actualValue,
		Comment:                 comment,
	}

	err := dbprovider.GetDBManager().UpdateChecklistItemInstance(item)
	if err != nil {
		ctx.ViewData("error", "Failed to update item")
	}

	ctx.Redirect(fmt.Sprintf("/sms_checklistInstance/show/%d", checklistInstanceID))
}

// POST: System-Checkliste instanziieren
func GenerateChecklistInstanceForSystem(ctx iris.Context) {
	systemID, _ := strconv.Atoi(ctx.Params().Get("system_id"))
	templateID, _ := strconv.Atoi(ctx.PostValue("ChecklistTemplateID"))

	includeDevice := ctx.PostValue("IncludeDeviceItems") == "on"
	versionStrategy := ctx.PostValueDefault("VersionStrategy", "all") // all|exact|wildcard
	versionPattern  := ctx.PostValue("VersionPattern")
	if versionStrategy != "wildcard" { versionPattern = "" }

	_, err := dbprovider.GetDBManager().AddChecklistInstanceForSystem(
		templateID, systemID, "system",
		includeDevice, versionStrategy, versionPattern,
	)
	if err != nil {
		ctx.ViewData("error", fmt.Sprintf("System-Checklist konnte nicht erzeugt werden: %v", err))
		system := dbprovider.GetDBManager().GetSMSSystemInfo(systemID)
		ctx.ViewData("system", system)
		ctx.View("sms_showSystem.html")
		return
	}
	ctx.Redirect(fmt.Sprintf("/sms_systems/show/%d", systemID))
}

// POST /sms_checklistTemplate/{id}/docasset/upload
// form-data: kind(cover|footer), mime(html|md), file (required)
func UploadChecklistTemplateDocAsset(ctx iris.Context) {
	templateID, err := ctx.Params().GetInt("id")
	if err != nil || templateID <= 0 {
		ctx.StatusCode(iris.StatusBadRequest)
		ctx.WriteString("invalid template id")
		return
	}

	kind := strings.ToLower(ctx.PostValue("kind"))
	// allow: cover | header | footer
	if kind != "cover" && kind != "header" && kind != "footer" {
		ctx.StatusCode(iris.StatusBadRequest)
		ctx.WriteString("kind must be cover, header or footer")
		return
	}
	mime := strings.ToLower(ctx.PostValue("mime"))
	if mime != "html" && mime != "md" && mime != "" {
		ctx.StatusCode(iris.StatusBadRequest)
		ctx.WriteString("mime must be html or md")
		return
	}
	if mime == "" { mime = "html" }

	file, fileHeader, err := ctx.FormFile("file")
	if err != nil {
		ctx.StatusCode(iris.StatusBadRequest)
		ctx.WriteString("file required")
		return
	}
	defer file.Close()

	// Dateiname bestimmen
	ext := ".html"
	if mime == "md" { ext = ".md" }
	// /var/app/docs/checklist-assets/{templateID}/cover.html
	dir := filepath.Join(docAssetBase, fmt.Sprintf("%d", templateID))
	if err := os.MkdirAll(dir, 0755); err != nil {
		ctx.StatusCode(iris.StatusInternalServerError)
		ctx.WriteString("cannot create asset dir")
		return
	}
	dstPath := filepath.Join(dir, kind+ext)

	// Datei speichern
	dst, err := os.Create(dstPath)
	if err != nil {
		ctx.StatusCode(iris.StatusInternalServerError)
		ctx.WriteString("cannot write asset file")
		return
	}
	defer dst.Close()
	if _, err := io.Copy(dst, file); err != nil {
		ctx.StatusCode(iris.StatusInternalServerError)
		ctx.WriteString("cannot save asset file")
		return
	}

	// DB upsert
	if err := dbprovider.GetDBManager().SaveChecklistTemplateDocAssetFile(templateID, kind, mime, dstPath); err != nil {
		ctx.StatusCode(iris.StatusInternalServerError)
		ctx.WriteString("cannot persist asset in db: " + err.Error())
		return
	}

	log.Printf("📄 uploaded %s (%s) for template %d: %s (src=%s)", kind, mime, templateID, dstPath, fileHeader.Filename)
	ctx.Redirect(fmt.Sprintf("/sms_checklistTemplate/show/%d", templateID))
}

// GET /sms_checklistTemplate/{id}/docasset/delete/{kind}
func DeleteChecklistTemplateDocAsset(ctx iris.Context) {
	templateID, err := ctx.Params().GetInt("id")
	if err != nil || templateID <= 0 {
		ctx.StatusCode(iris.StatusBadRequest)
		ctx.WriteString("invalid template id")
		return
	}
	kind := strings.ToLower(ctx.Params().GetStringDefault("kind", ""))
	if kind != "cover" && kind != "footer" {
		ctx.StatusCode(iris.StatusBadRequest)
		ctx.WriteString("kind must be cover or footer")
		return
	}

	// DB löschen
	if err := dbprovider.GetDBManager().DeleteChecklistTemplateDocAsset(templateID, kind); err != nil {
		ctx.StatusCode(iris.StatusInternalServerError)
		ctx.WriteString("cannot delete doc asset: " + err.Error())
		return
	}

	// Datei optional entfernen (wir löschen beide möglichen Endungen)
	dir := filepath.Join(docAssetBase, fmt.Sprintf("%d", templateID))
	_ = os.Remove(filepath.Join(dir, kind+".html"))
	_ = os.Remove(filepath.Join(dir, kind+".md"))

	log.Printf("🗑️ removed %s doc asset for template %d", kind, templateID)
	ctx.Redirect(fmt.Sprintf("/sms_checklistTemplate/show/%d", templateID))
}

// GET /sms_checklistInstance/print/{id:int}
func PrintChecklistInstance(ctx iris.Context) {
	id, err := ctx.Params().GetInt("id")
	if err != nil || id <= 0 {
		ctx.StatusCode(iris.StatusBadRequest)
		_, _ = ctx.WriteString("invalid checklist instance id")
		return
	}

	inst := dbprovider.GetDBManager().GetChecklistInstanceByID(id)
	if inst == nil {
		ctx.StatusCode(iris.StatusNotFound)
		_, _ = ctx.WriteString("instance not found")
		return
	}
	items := dbprovider.GetDBManager().GetChecklistItemInstancesWithDefinition(id)

	// Falls TemplateName hier nicht gesetzt ist: (optional) aus Template-Tabelle nachladen
	if inst.TemplateName == "" {
		if name := dbprovider.GetDBManager().GetChecklistTemplateName(inst.ChecklistTemplateID); name != "" {
			inst.TemplateName = name
		}
	}

	// absolute Logo-URL (wkhtmltopdf ruft via HTTP auf)
	scheme := "http"
	if ctx.Request().TLS != nil { scheme = "https" }
	logoURL := fmt.Sprintf("%s://%s/static/img/logo.png", scheme, ctx.Host())

	// Kein globales Layout erzwingen:
	ctx.ViewLayout(iris.NoLayout)

	// „now“ als String, keine |date-Filter nötig
	nowStr := time.Now().Format("2006-01-02 15:04")

	dbprovider.GetDBManager().PostRenderChecklistItems(inst, items)

	ctx.ViewData("instance", inst)
	ctx.ViewData("items", items)
	ctx.ViewData("now", nowStr)
	ctx.ViewData("logoPath", logoURL)

	if err := ctx.View("sms_exportChecklistInstance.html"); err != nil {
		// 500 loggen + Fehltext ausgeben, damit du im Browser/wkhtmltopdf siehst, WAS kaputt ist
		ctx.StatusCode(iris.StatusInternalServerError)
		log.Printf("print view render error: %v", err)
		_, _ = ctx.WriteString("render error: " + err.Error())
		return
	}
}

// GET /sms_checklistInstance/export/{id:int}?fmt=pdf
func ExportChecklistInstance(ctx iris.Context) {
	id, err := ctx.Params().GetInt("id")
	if err != nil || id <= 0 {
		ctx.StatusCode(iris.StatusBadRequest)
		_, _ = ctx.WriteString("invalid checklist instance id")
		return
	}

	inst := dbprovider.GetDBManager().GetChecklistInstanceByID(id)
	if inst == nil {
		ctx.StatusCode(iris.StatusNotFound)
		_, _ = ctx.WriteString("instance not found")
		return
	}

	scheme := "http"
	if ctx.Request().TLS != nil {
		scheme = "https"
	}
	host := ctx.Host()
	printURL := fmt.Sprintf("%s://%s/sms_checklistInstance/print/%d", scheme, host, id)

	outFile := filepath.Join(os.TempDir(), fmt.Sprintf("checklist_%d.pdf", id))

	args := []string{
		"--enable-local-file-access",
		"--encoding", "utf-8",
		"--page-size", "A4",
		"--margin-top", "12mm",
		"--margin-bottom", "12mm",
	}

	// Optional: Cover / Header / Footer URLs holen
	coverURL, headerURL, footerURL := dbprovider.GetDBManager().GetDocAssetURLs(docAssetBase, inst.ChecklistTemplateID, scheme, host)
	if coverURL != "" {
		args = append(args, "cover", coverURL)
	}

	// Hauptinhalt
	args = append(args, printURL)

	if headerURL != "" {
		args = append(args, "--header-html", headerURL, "--header-spacing", "5")
	}
	if footerURL != "" {
		args = append(args, "--footer-html", footerURL, "--footer-spacing", "5")
	}

	// Ausgabe-Datei
	args = append(args, outFile)

	cmd := exec.Command("wkhtmltopdf", args...)
	if out, err := cmd.CombinedOutput(); err != nil {
		log.Printf("wkhtmltopdf error: %s", string(out))
		ctx.StatusCode(iris.StatusInternalServerError)
		_, _ = ctx.WriteString("PDF generation failed")
		return
	}

	ctx.Header("Content-Type", "application/pdf")
	ctx.Header("Content-Disposition", fmt.Sprintf(`attachment; filename="checklist_%d.pdf"`, id))
	http.ServeFile(ctx.ResponseWriter(), ctx.Request(), outFile)
}

func sendDocAsset(ctx iris.Context, pathHTML, pathMD, downloadName string) {
	if st, err := os.Stat(pathHTML); err == nil && !st.IsDir() {
		ctx.ContentType("text/html")
		ctx.SendFile(pathHTML, downloadName)
		return
	}
	if st, err := os.Stat(pathMD); err == nil && !st.IsDir() {
		b, err := os.ReadFile(pathMD)
		if err != nil {
			ctx.StatusCode(iris.StatusInternalServerError)
			_, _ = ctx.WriteString("cannot read md")
			return
		}
		// super-einfache Umwandlung; nimm gern eine Markdown-Lib (goldmark) falls du willst
		html := "<!doctype html><meta charset=\"utf-8\"><body><pre>" +
			htmlEscape(string(b)) + "</pre></body>"
		ctx.ContentType("text/html")
		_, _ = ctx.WriteString(html)
		return
	}
	ctx.StatusCode(iris.StatusNotFound)
	_, _ = ctx.WriteString("no asset")
}

func DocAssetCover(ctx iris.Context) {
	tid, _ := ctx.Params().GetInt("template_id")
	base := filepath.Join(docAssetBase, fmt.Sprintf("%d", tid))
	sendDocAsset(ctx,
		filepath.Join(base, "cover.html"),
		filepath.Join(base, "cover.md"),
		"cover.html",
	)
}

func DocAssetHeader(ctx iris.Context) {
	tid, _ := ctx.Params().GetInt("template_id")
	base := filepath.Join(docAssetBase, fmt.Sprintf("%d", tid))
	sendDocAsset(ctx,
		filepath.Join(base, "header.html"),
		filepath.Join(base, "header.md"),
		"header.html",
	)
}

func DocAssetFooter(ctx iris.Context) {
	tid, _ := ctx.Params().GetInt("template_id")
	base := filepath.Join(docAssetBase, fmt.Sprintf("%d", tid))
	sendDocAsset(ctx,
		filepath.Join(base, "footer.html"),
		filepath.Join(base, "footer.md"),
		"footer.html",
	)
}

func htmlEscape(s string) string {
	r := strings.NewReplacer(
		"&", "&amp;", "<", "&lt;", ">", "&gt;", `"`, "&#34;", "'", "&#39;",
	)
	return r.Replace(s)
}

// in dbprovider or routes utils – your choice
type DocAssetInfo struct {
	HasCover  bool
	HasHeader bool
	HasFooter bool
	CoverURL  string
	HeaderURL string
	FooterURL string
	DirPath   string
}

func GetDocAssetInfoForTemplate(templateID int, scheme, host string) DocAssetInfo {
	base := filepath.Join(docAssetBase, fmt.Sprintf("%d", templateID))
	info := DocAssetInfo{
		DirPath: base,
		CoverURL:  fmt.Sprintf("%s://%s/docassets/cover/%d",  scheme, host, templateID),
		HeaderURL: fmt.Sprintf("%s://%s/docassets/header/%d", scheme, host, templateID),
		FooterURL: fmt.Sprintf("%s://%s/docassets/footer/%d", scheme, host, templateID),
	}
	// a file counts if either HTML or MD exists
	if _, err := os.Stat(filepath.Join(base, "cover.html")); err == nil { info.HasCover = true }
	if _, err := os.Stat(filepath.Join(base, "cover.md"));   err == nil { info.HasCover = true }

	if _, err := os.Stat(filepath.Join(base, "header.html")); err == nil { info.HasHeader = true }
	if _, err := os.Stat(filepath.Join(base, "header.md"));   err == nil { info.HasHeader = true }

	if _, err := os.Stat(filepath.Join(base, "footer.html")); err == nil { info.HasFooter = true }
	if _, err := os.Stat(filepath.Join(base, "footer.md"));   err == nil { info.HasFooter = true }

	return info
}
