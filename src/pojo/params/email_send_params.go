package params

import "github.com/bingfenglai/gt/errors"

// 发送邮件参数
type EmailSimpleSendParams struct {
	// 接收者
	Receivers []string
	//// 密送
	//Bcc []string
	//
	//// 抄送
	//Cc []string

	// 邮件主题
	Subject string

	Text        []byte // Plaintext message (optional)
	HTML        []byte // Html message (optional)
	//Sender      string // override From as SMTP envelope sender (optional)
	//Headers     textproto.MIMEHeader  // 协议头
	//Attachments []*Attachment    // 附件
	//ReadReceipt []string
}

func (p EmailSimpleSendParams) Check() error {

	if p.HTML == nil && p.Text == nil {
		return errors.ErrEmailContentIsNull
	}

	if p.Receivers == nil || len(p.Receivers) == 0 {
		return errors.ErrEmailContentIsNull
	}

	if p.Subject == "" {
		return errors.ErrEmailContentIsNull
	}

	return nil


}
