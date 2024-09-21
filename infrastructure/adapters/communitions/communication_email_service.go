package communitions

import "fmt"

type CommunicationEmailService struct {
}

func NewCommunicationEmailService() *CommunicationEmailService {
	return &CommunicationEmailService{}
}

func (c CommunicationEmailService) SendOpt(emailAddress, name, otp, systemName string) error {

	var subject = fmt.Sprintf("%v One-Time Password", systemName)
	var senderName = fmt.Sprintf("%v Health Service Authentication", systemName)
	var org = systemName
	var body = fmt.Sprintf(` 
		<h2>Hello %v</h2>
		<p>Your One-Time Password (OTP) for reset is: %v.</p>
		<p>Please use this OTP to complete your action within the next 2 hours.</p>
		<p />
		<p> Thank you</p>
		<p> %v Service</>
		`,
		name, otp, systemName)
	var myRequest EmailRequest
	var from = ""
	myRequest.Body = body
	myRequest.Subject = subject
	myRequest.To = emailAddress
	myRequest.From = from
	myRequest.SenderName = senderName
	myRequest.Org = org
	err := SendSampleEmail(myRequest)
	return err
}
