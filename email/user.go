package email

// SendVerifyCode send verify code to user
func SendVerifyCode(sendTo string, vc string) error {
	var eData emailData
	eData.Subject = "[Application]: Your verify code Info"
	eData.Tittle = "please check your verify code "
	eData.SendTo = sendTo
	eData.Content = "<h1>Your verify code ：</h1>\n"

	eData.Content = vc + "<br>"

	return sendEmail(eData)
}
