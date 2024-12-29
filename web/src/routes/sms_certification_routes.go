/**
 * Author:    Admiral Helmut
 * Created:   12.06.2019
 *
 * (C)
 **/

package routes

import (
	"github.com/efi4st/efi4st/dbprovider"
	_ "github.com/go-sql-driver/mysql"
	"github.com/kataras/iris/v12"
	"strconv"
)

func SMSCertification(ctx iris.Context) {

	certifications := dbprovider.GetDBManager().GetSMSCertification()
	ctx.ViewData("error", "")

	if len(certifications) < 1 {
		ctx.ViewData("error", "Error: No certifications available. Add one!")
	}
	ctx.ViewData("certificationList", certifications)
	ctx.View("sms_certifications.html")
}

// GET
func CreateSMSCertification(ctx iris.Context) {

	ctx.View("sms_createCertification.html")
}

// POST
func AddSMSCertification(ctx iris.Context) {

	certificationName := ctx.PostValue("CertificationName")
	description := ctx.PostValue("Description")

	err := dbprovider.GetDBManager().AddSMSCertification(certificationName, description)

	ctx.ViewData("error", "")
	if err !=nil {
		ctx.ViewData("error", "Error: Not able to add certification!")
	}
	certifications := dbprovider.GetDBManager().GetSMSCertification()
	ctx.ViewData("certificationList", certifications)
	ctx.View("sms_certifications.html")
}

func ShowSMSCertification(ctx iris.Context) {

	id := ctx.Params().Get("id")

	i, err := strconv.Atoi(id)

	ctx.ViewData("error", "")
	if err !=nil {
		ctx.ViewData("error", "Error: Error parsing certification Id!")
	}

	certification := dbprovider.GetDBManager().GetSMSCertificationInfo(i)
	systemHasCertification, err := dbprovider.GetDBManager().GetSystemsForCertification(i)
	if err != nil {
		ctx.ViewData("error", "Error: Error matching Systems to certificate!")
	}

	ctx.ViewData("systemHasCertification", systemHasCertification)
	ctx.ViewData("certification", certification)
	ctx.View("sms_showCertification.html")
}

func RemoveSMSCertification(ctx iris.Context) {

	id := ctx.Params().Get("id")

	i, err := strconv.Atoi(id)
	err = dbprovider.GetDBManager().RemoveSMSCertification(i)

	ctx.ViewData("error", "")
	if err !=nil {
		ctx.ViewData("error", "Error: Error removing certification!")
	}

	certifications := dbprovider.GetDBManager().GetSMSCertification()

	ctx.ViewData("certificationList", certifications)
	ctx.View("sms_certifications.html")
}