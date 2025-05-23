/*
 * @license
 * Copyright 2025 Dynatrace LLC
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 * http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package download

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOnlyOptions_IsSingleOption(t *testing.T) {
	t.Run("Should return true if set and the only one", func(t *testing.T) {
		onlyOption := OnlyOptions{OnlySettingsFlag: true}

		got := onlyOption.IsSingleOption(OnlySettingsFlag)
		assert.True(t, got)
	})

	t.Run("Should return false if set and not the only one", func(t *testing.T) {
		onlyOption := OnlyOptions{OnlySettingsFlag: true, OnlyApisFlag: true}

		got := onlyOption.IsSingleOption(OnlySettingsFlag)
		assert.False(t, got)
	})

	t.Run("Should return false if not set", func(t *testing.T) {
		onlyOption := OnlyOptions{OnlyApisFlag: true}

		got := onlyOption.IsSingleOption(OnlySettingsFlag)
		assert.False(t, got)
	})

	t.Run("Should return false if set but false", func(t *testing.T) {
		onlyOption := OnlyOptions{OnlySettingsFlag: false}

		got := onlyOption.IsSingleOption(OnlySettingsFlag)
		assert.False(t, got)
	})
}

func TestOnlyOptions_OnlyCount(t *testing.T) {
	t.Run("Should return the amount of enabled flags", func(t *testing.T) {
		onlyOption := OnlyOptions{OnlySettingsFlag: false, OnlyApisFlag: true, OnlySegmentsFlag: true}

		got := onlyOption.OnlyCount()
		assert.Equal(t, got, 2)
	})
}

func TestOnlyOptions_ShouldDownload(t *testing.T) {
	t.Run("Should return true if no flags", func(t *testing.T) {
		onlyOption := OnlyOptions{}

		got := onlyOption.ShouldDownload(OnlySettingsFlag)
		assert.True(t, got)
	})

	t.Run("Should return true if no true flags are set", func(t *testing.T) {
		onlyOption := OnlyOptions{OnlySettingsFlag: false, OnlyApisFlag: false, OnlySegmentsFlag: false}

		got := onlyOption.ShouldDownload(OnlySettingsFlag)
		assert.True(t, got)
	})

	t.Run("Should return true if the only flag is set", func(t *testing.T) {
		onlyOption := OnlyOptions{OnlySettingsFlag: true, OnlyApisFlag: false, OnlySegmentsFlag: true}

		got := onlyOption.ShouldDownload(OnlySettingsFlag)
		assert.True(t, got)
	})

	t.Run("Should return false", func(t *testing.T) {
		onlyOption := OnlyOptions{OnlySettingsFlag: true, OnlyApisFlag: false, OnlySegmentsFlag: true}

		got := onlyOption.ShouldDownload(OnlyApisFlag)
		assert.False(t, got)
	})
}
