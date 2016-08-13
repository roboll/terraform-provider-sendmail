package main

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/hashicorp/terraform/helper/schema"
	"github.com/satori/go.uuid"
)

func resourceSend() *schema.Resource {
	return &schema.Resource{
		Create: resourceSendCreate,
		Read:   noop,
		Update: resourceSendCreate,
		Delete: resourceSendDelete,

		Schema: map[string]*schema.Schema{
			"from": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: false,
			},
			"to": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: false,
			},

			"subject": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},

			"body": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
		},
	}
}

const (
	sendmail = "/usr/sbin/sendmail"
	template = `To: %s
From: %s
Subject: %s

%s`
)

func resourceSendCreate(d *schema.ResourceData, meta interface{}) error {
	from := d.Get("from").(string)
	to := d.Get("to").(string)
	subject := d.Get("subject").(string)
	body := d.Get("body").(string)

	if err := send(from, to, subject, body); err != nil {
		return err
	}

	d.SetId(uuid.NewV4().String())
	return nil
}

func resourceSendDelete(d *schema.ResourceData, meta interface{}) error {
	d.SetId("")
	return nil
}

func noop(d *schema.ResourceData, meta interface{}) error {
	return nil
}

func send(from, to, subject, body string) (err error) {
	cmd := exec.Command(sendmail, "-t")

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	pw, err := cmd.StdinPipe()
	if err != nil {
		return
	}

	err = cmd.Start()
	if err != nil {
		return
	}

	var errs [3]error
	_, errs[0] = pw.Write([]byte(fmt.Sprintf(template, to, from, subject, body)))
	errs[1] = pw.Close()
	errs[2] = cmd.Wait()
	for _, err = range errs {
		if err != nil {
			return
		}
	}
	return
}
