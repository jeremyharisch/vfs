/*
 * Copyright 2020 Mandelsoft. All rights reserved.
 *  This file is licensed under the Apache Software License, v. 2 except as noted
 *  otherwise in the LICENSE file
 *
 *  Licensed under the Apache License, Version 2.0 (the "License");
 *  you may not use this file except in compliance with the License.
 *  You may obtain a copy of the License at
 *
 *       http://www.apache.org/licenses/LICENSE-2.0
 *
 *  Unless required by applicable law or agreed to in writing, software
 *  distributed under the License is distributed on an "AS IS" BASIS,
 *  WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 *  See the License for the specific language governing permissions and
 *  limitations under the License.
 */

package memoryfs

import (
	"os"
	"sort"
)

type DirectoryEntries map[string]*fileData

func (m DirectoryEntries) Len() int                     { return len(m) }
func (m DirectoryEntries) Add(name string, f *fileData) { m[name] = f }
func (m DirectoryEntries) Remove(name string)           { delete(m, name) }
func (m DirectoryEntries) Files() (files []os.FileInfo) {
	for n, f := range m {
		files = append(files, newFileInfo(n, f))
	}
	sort.Sort(filesSorter(files))
	return files
}

type filesSorter []os.FileInfo

func (s filesSorter) Len() int           { return len(s) }
func (s filesSorter) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }
func (s filesSorter) Less(i, j int) bool { return s[i].Name() < s[j].Name() }

func (m DirectoryEntries) Names() (names []string) {
	for x := range m {
		names = append(names, x)
	}
	return names
}