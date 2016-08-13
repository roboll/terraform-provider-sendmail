# terraform-provider-sendmail

A terraform provider for sending mail using the local `/usr/sbin/sendmail` utility.

## usage

```
resource sendmail_send email {
  from = "someone@example.com"
  to = "otherone@example.com"
  subject = "A Terraform Email"
  body = <<EMAIL
Hello, this is an email from terraform.
EMAIL
}
```

## get it

`go get github.com/roboll/terraform-provider-sendmail`

_or_

`curl -L -o /usr/local/bin/terraform-provider-sendmail https://github.com/roboll/terraform-provider-sendmail/releases/download/{VERSION}/terraform-provider-sendmail_{OS}_{ARCH}`

## development

[govendor](https://github.com/kardianos/govendor) for vendoring
