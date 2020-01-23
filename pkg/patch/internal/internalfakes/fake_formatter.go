/*
Copyright 2020 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by counterfeiter. DO NOT EDIT.
package internalfakes

import (
	"sync"

	"k8s.io/release/pkg/patch"
)

type FakeFormatter struct {
	MarkdownToHTMLStub        func(string, string) (string, error)
	markdownToHTMLMutex       sync.RWMutex
	markdownToHTMLArgsForCall []struct {
		arg1 string
		arg2 string
	}
	markdownToHTMLReturns struct {
		result1 string
		result2 error
	}
	markdownToHTMLReturnsOnCall map[int]struct {
		result1 string
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeFormatter) MarkdownToHTML(arg1 string, arg2 string) (string, error) {
	fake.markdownToHTMLMutex.Lock()
	ret, specificReturn := fake.markdownToHTMLReturnsOnCall[len(fake.markdownToHTMLArgsForCall)]
	fake.markdownToHTMLArgsForCall = append(fake.markdownToHTMLArgsForCall, struct {
		arg1 string
		arg2 string
	}{arg1, arg2})
	fake.recordInvocation("MarkdownToHTML", []interface{}{arg1, arg2})
	fake.markdownToHTMLMutex.Unlock()
	if fake.MarkdownToHTMLStub != nil {
		return fake.MarkdownToHTMLStub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.markdownToHTMLReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeFormatter) MarkdownToHTMLCallCount() int {
	fake.markdownToHTMLMutex.RLock()
	defer fake.markdownToHTMLMutex.RUnlock()
	return len(fake.markdownToHTMLArgsForCall)
}

func (fake *FakeFormatter) MarkdownToHTMLCalls(stub func(string, string) (string, error)) {
	fake.markdownToHTMLMutex.Lock()
	defer fake.markdownToHTMLMutex.Unlock()
	fake.MarkdownToHTMLStub = stub
}

func (fake *FakeFormatter) MarkdownToHTMLArgsForCall(i int) (string, string) {
	fake.markdownToHTMLMutex.RLock()
	defer fake.markdownToHTMLMutex.RUnlock()
	argsForCall := fake.markdownToHTMLArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeFormatter) MarkdownToHTMLReturns(result1 string, result2 error) {
	fake.markdownToHTMLMutex.Lock()
	defer fake.markdownToHTMLMutex.Unlock()
	fake.MarkdownToHTMLStub = nil
	fake.markdownToHTMLReturns = struct {
		result1 string
		result2 error
	}{result1, result2}
}

func (fake *FakeFormatter) MarkdownToHTMLReturnsOnCall(i int, result1 string, result2 error) {
	fake.markdownToHTMLMutex.Lock()
	defer fake.markdownToHTMLMutex.Unlock()
	fake.MarkdownToHTMLStub = nil
	if fake.markdownToHTMLReturnsOnCall == nil {
		fake.markdownToHTMLReturnsOnCall = make(map[int]struct {
			result1 string
			result2 error
		})
	}
	fake.markdownToHTMLReturnsOnCall[i] = struct {
		result1 string
		result2 error
	}{result1, result2}
}

func (fake *FakeFormatter) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.markdownToHTMLMutex.RLock()
	defer fake.markdownToHTMLMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeFormatter) recordInvocation(key string, args []interface{}) {
	fake.invocationsMutex.Lock()
	defer fake.invocationsMutex.Unlock()
	if fake.invocations == nil {
		fake.invocations = map[string][][]interface{}{}
	}
	if fake.invocations[key] == nil {
		fake.invocations[key] = [][]interface{}{}
	}
	fake.invocations[key] = append(fake.invocations[key], args)
}

var _ patch.Formatter = new(FakeFormatter)
