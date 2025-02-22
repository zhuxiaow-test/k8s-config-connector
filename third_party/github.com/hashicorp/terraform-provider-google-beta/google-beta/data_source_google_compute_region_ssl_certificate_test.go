// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
package google

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/acctest"
)

func TestAccDataSourceComputeRegionSslCertificate(t *testing.T) {
	t.Parallel()

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceComputeRegionSslCertificateConfig(acctest.RandString(t, 10)),
				Check: resource.ComposeTestCheckFunc(
					acctest.CheckDataSourceStateMatchesResourceStateWithIgnores(
						"data.google_compute_region_ssl_certificate.cert",
						"google_compute_region_ssl_certificate.foobar",
						map[string]struct{}{
							"private_key": {},
						},
					),
				),
			},
		},
	})
}

func testAccDataSourceComputeRegionSslCertificateConfig(certName string) string {
	return fmt.Sprintf(`
resource "google_compute_region_ssl_certificate" "foobar" {
  name        = "cert-test-%s"
  region      = "us-central1"
  description = "really descriptive"
  private_key = file("test-fixtures/ssl_cert/test.key")
  certificate = file("test-fixtures/ssl_cert/test.crt")
}

data "google_compute_region_ssl_certificate" "cert" {
  name = google_compute_region_ssl_certificate.foobar.name
}
`, certName)
}
