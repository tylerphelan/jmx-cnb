/*
 * Copyright 2018-2019 the original author or authors.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *      http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package main

import (
	"testing"

	"github.com/buildpack/libbuildpack/buildplan"
	"github.com/cloudfoundry/jmx-cnb/jmx"
	"github.com/cloudfoundry/jvm-application-cnb/jvmapplication"
	"github.com/cloudfoundry/libcfbuildpack/detect"
	"github.com/cloudfoundry/libcfbuildpack/test"
	. "github.com/onsi/gomega"
	"github.com/sclevine/spec"
	"github.com/sclevine/spec/report"
)

func TestDetect(t *testing.T) {
	spec.Run(t, "Detect", func(t *testing.T, _ spec.G, it spec.S) {

		g := NewGomegaWithT(t)

		var f *test.DetectFactory

		it.Before(func() {
			f = test.NewDetectFactory(t)
		})

		it("fails without jvm-application", func() {
			defer test.ReplaceEnv(t, "BP_JMX", "")()

			g.Expect(d(f.Detect)).To(Equal(detect.FailStatusCode))
		})

		it("fails without BP_JMX", func() {
			f.AddBuildPlan(jvmapplication.Dependency, buildplan.Dependency{})

			g.Expect(d(f.Detect)).To(Equal(detect.FailStatusCode))
		})

		it("passes with jvm-application and BP_JMX", func() {
			f.AddBuildPlan(jvmapplication.Dependency, buildplan.Dependency{})
			defer test.ReplaceEnv(t, "BP_JMX", "")()

			g.Expect(d(f.Detect)).To(Equal(detect.PassStatusCode))
			g.Expect(f.Output).To(Equal(buildplan.BuildPlan{
				jmx.Dependency: buildplan.Dependency{},
			}))
		})
	}, spec.Report(report.Terminal{}))
}
