// Code generated by easyjson for marshaling/unmarshaling. DO NOT EDIT.

package hyper

import (
	json "encoding/json"
	easyjson "github.com/mailru/easyjson"
	jlexer "github.com/mailru/easyjson/jlexer"
	jwriter "github.com/mailru/easyjson/jwriter"
)

// suppress unused package warning
var (
	_ *json.RawMessage
	_ *jlexer.Lexer
	_ *jwriter.Writer
	_ easyjson.Marshaler
)

func easyjsonD2b7633eDecodeGithubComHyperSolutionsHyperSdkGo(in *jlexer.Lexer, out *kasadaPayloadOutput) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeFieldName(false)
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "headers":
			(out.Headers).UnmarshalEasyJSON(in)
		case "payload":
			out.Payload = string(in.String())
		case "error":
			out.Error = string(in.String())
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjsonD2b7633eEncodeGithubComHyperSolutionsHyperSdkGo(out *jwriter.Writer, in kasadaPayloadOutput) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"headers\":"
		out.RawString(prefix[1:])
		(in.Headers).MarshalEasyJSON(out)
	}
	{
		const prefix string = ",\"payload\":"
		out.RawString(prefix)
		out.String(string(in.Payload))
	}
	{
		const prefix string = ",\"error\":"
		out.RawString(prefix)
		out.String(string(in.Error))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v kasadaPayloadOutput) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonD2b7633eEncodeGithubComHyperSolutionsHyperSdkGo(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v kasadaPayloadOutput) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonD2b7633eEncodeGithubComHyperSolutionsHyperSdkGo(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *kasadaPayloadOutput) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonD2b7633eDecodeGithubComHyperSolutionsHyperSdkGo(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *kasadaPayloadOutput) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonD2b7633eDecodeGithubComHyperSolutionsHyperSdkGo(l, v)
}
func easyjsonD2b7633eDecodeGithubComHyperSolutionsHyperSdkGo1(in *jlexer.Lexer, out *apiResponse) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeFieldName(false)
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "payload":
			out.Payload = string(in.String())
		case "error":
			out.Error = string(in.String())
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjsonD2b7633eEncodeGithubComHyperSolutionsHyperSdkGo1(out *jwriter.Writer, in apiResponse) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"payload\":"
		out.RawString(prefix[1:])
		out.String(string(in.Payload))
	}
	{
		const prefix string = ",\"error\":"
		out.RawString(prefix)
		out.String(string(in.Error))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v apiResponse) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonD2b7633eEncodeGithubComHyperSolutionsHyperSdkGo1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v apiResponse) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonD2b7633eEncodeGithubComHyperSolutionsHyperSdkGo1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *apiResponse) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonD2b7633eDecodeGithubComHyperSolutionsHyperSdkGo1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *apiResponse) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonD2b7633eDecodeGithubComHyperSolutionsHyperSdkGo1(l, v)
}
func easyjsonD2b7633eDecodeGithubComHyperSolutionsHyperSdkGo2(in *jlexer.Lexer, out *UtmvcInput) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeFieldName(false)
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "userAgent":
			out.UserAgent = string(in.String())
		case "sessionIds":
			if in.IsNull() {
				in.Skip()
				out.SessionIds = nil
			} else {
				in.Delim('[')
				if out.SessionIds == nil {
					if !in.IsDelim(']') {
						out.SessionIds = make([]string, 0, 4)
					} else {
						out.SessionIds = []string{}
					}
				} else {
					out.SessionIds = (out.SessionIds)[:0]
				}
				for !in.IsDelim(']') {
					var v1 string
					v1 = string(in.String())
					out.SessionIds = append(out.SessionIds, v1)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "script":
			out.Script = string(in.String())
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjsonD2b7633eEncodeGithubComHyperSolutionsHyperSdkGo2(out *jwriter.Writer, in UtmvcInput) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"userAgent\":"
		out.RawString(prefix[1:])
		out.String(string(in.UserAgent))
	}
	{
		const prefix string = ",\"sessionIds\":"
		out.RawString(prefix)
		if in.SessionIds == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v2, v3 := range in.SessionIds {
				if v2 > 0 {
					out.RawByte(',')
				}
				out.String(string(v3))
			}
			out.RawByte(']')
		}
	}
	{
		const prefix string = ",\"script\":"
		out.RawString(prefix)
		out.String(string(in.Script))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v UtmvcInput) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonD2b7633eEncodeGithubComHyperSolutionsHyperSdkGo2(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v UtmvcInput) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonD2b7633eEncodeGithubComHyperSolutionsHyperSdkGo2(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *UtmvcInput) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonD2b7633eDecodeGithubComHyperSolutionsHyperSdkGo2(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *UtmvcInput) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonD2b7633eDecodeGithubComHyperSolutionsHyperSdkGo2(l, v)
}
func easyjsonD2b7633eDecodeGithubComHyperSolutionsHyperSdkGo3(in *jlexer.Lexer, out *SensorInput) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeFieldName(false)
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "abck":
			out.Abck = string(in.String())
		case "bmsz":
			out.Bmsz = string(in.String())
		case "version":
			out.Version = string(in.String())
		case "pageUrl":
			out.PageUrl = string(in.String())
		case "userAgent":
			out.UserAgent = string(in.String())
		case "scriptHash":
			out.ScriptHash = string(in.String())
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjsonD2b7633eEncodeGithubComHyperSolutionsHyperSdkGo3(out *jwriter.Writer, in SensorInput) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"abck\":"
		out.RawString(prefix[1:])
		out.String(string(in.Abck))
	}
	{
		const prefix string = ",\"bmsz\":"
		out.RawString(prefix)
		out.String(string(in.Bmsz))
	}
	{
		const prefix string = ",\"version\":"
		out.RawString(prefix)
		out.String(string(in.Version))
	}
	{
		const prefix string = ",\"pageUrl\":"
		out.RawString(prefix)
		out.String(string(in.PageUrl))
	}
	{
		const prefix string = ",\"userAgent\":"
		out.RawString(prefix)
		out.String(string(in.UserAgent))
	}
	{
		const prefix string = ",\"scriptHash\":"
		out.RawString(prefix)
		out.String(string(in.ScriptHash))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v SensorInput) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonD2b7633eEncodeGithubComHyperSolutionsHyperSdkGo3(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v SensorInput) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonD2b7633eEncodeGithubComHyperSolutionsHyperSdkGo3(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *SensorInput) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonD2b7633eDecodeGithubComHyperSolutionsHyperSdkGo3(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *SensorInput) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonD2b7633eDecodeGithubComHyperSolutionsHyperSdkGo3(l, v)
}
func easyjsonD2b7633eDecodeGithubComHyperSolutionsHyperSdkGo4(in *jlexer.Lexer, out *PixelInput) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeFieldName(false)
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "userAgent":
			out.UserAgent = string(in.String())
		case "htmlVar":
			out.HTMLVar = string(in.String())
		case "scriptVar":
			out.ScriptVar = string(in.String())
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjsonD2b7633eEncodeGithubComHyperSolutionsHyperSdkGo4(out *jwriter.Writer, in PixelInput) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"userAgent\":"
		out.RawString(prefix[1:])
		out.String(string(in.UserAgent))
	}
	{
		const prefix string = ",\"htmlVar\":"
		out.RawString(prefix)
		out.String(string(in.HTMLVar))
	}
	{
		const prefix string = ",\"scriptVar\":"
		out.RawString(prefix)
		out.String(string(in.ScriptVar))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v PixelInput) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonD2b7633eEncodeGithubComHyperSolutionsHyperSdkGo4(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v PixelInput) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonD2b7633eEncodeGithubComHyperSolutionsHyperSdkGo4(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *PixelInput) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonD2b7633eDecodeGithubComHyperSolutionsHyperSdkGo4(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *PixelInput) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonD2b7633eDecodeGithubComHyperSolutionsHyperSdkGo4(l, v)
}
func easyjsonD2b7633eDecodeGithubComHyperSolutionsHyperSdkGo5(in *jlexer.Lexer, out *KasadaPowInput) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeFieldName(false)
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "st":
			out.St = int(in.Int())
		case "workTime":
			if in.IsNull() {
				in.Skip()
				out.WorkTime = nil
			} else {
				if out.WorkTime == nil {
					out.WorkTime = new(int)
				}
				*out.WorkTime = int(in.Int())
			}
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjsonD2b7633eEncodeGithubComHyperSolutionsHyperSdkGo5(out *jwriter.Writer, in KasadaPowInput) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"st\":"
		out.RawString(prefix[1:])
		out.Int(int(in.St))
	}
	if in.WorkTime != nil {
		const prefix string = ",\"workTime\":"
		out.RawString(prefix)
		out.Int(int(*in.WorkTime))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v KasadaPowInput) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonD2b7633eEncodeGithubComHyperSolutionsHyperSdkGo5(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v KasadaPowInput) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonD2b7633eEncodeGithubComHyperSolutionsHyperSdkGo5(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *KasadaPowInput) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonD2b7633eDecodeGithubComHyperSolutionsHyperSdkGo5(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *KasadaPowInput) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonD2b7633eDecodeGithubComHyperSolutionsHyperSdkGo5(l, v)
}
func easyjsonD2b7633eDecodeGithubComHyperSolutionsHyperSdkGo6(in *jlexer.Lexer, out *KasadaPayloadInput) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeFieldName(false)
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "userAgent":
			out.UserAgent = string(in.String())
		case "ipsLink":
			out.IpsLink = string(in.String())
		case "script":
			out.Script = string(in.String())
		case "language":
			out.Language = string(in.String())
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjsonD2b7633eEncodeGithubComHyperSolutionsHyperSdkGo6(out *jwriter.Writer, in KasadaPayloadInput) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"userAgent\":"
		out.RawString(prefix[1:])
		out.String(string(in.UserAgent))
	}
	{
		const prefix string = ",\"ipsLink\":"
		out.RawString(prefix)
		out.String(string(in.IpsLink))
	}
	{
		const prefix string = ",\"script\":"
		out.RawString(prefix)
		out.String(string(in.Script))
	}
	if in.Language != "" {
		const prefix string = ",\"language\":"
		out.RawString(prefix)
		out.String(string(in.Language))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v KasadaPayloadInput) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonD2b7633eEncodeGithubComHyperSolutionsHyperSdkGo6(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v KasadaPayloadInput) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonD2b7633eEncodeGithubComHyperSolutionsHyperSdkGo6(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *KasadaPayloadInput) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonD2b7633eDecodeGithubComHyperSolutionsHyperSdkGo6(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *KasadaPayloadInput) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonD2b7633eDecodeGithubComHyperSolutionsHyperSdkGo6(l, v)
}
func easyjsonD2b7633eDecodeGithubComHyperSolutionsHyperSdkGo7(in *jlexer.Lexer, out *KasadaHeaders) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeFieldName(false)
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "x-kpsdk-ct":
			out.XKpsdkCt = string(in.String())
		case "x-kpsdk-dt":
			out.XKpsdkDt = string(in.String())
		case "x-kpsdk-v":
			out.XKpsdkV = string(in.String())
		case "x-kpsdk-r":
			out.XKpsdkR = string(in.String())
		case "x-kpsdk-dv":
			out.XKpsdkDv = string(in.String())
		case "x-kpsdk-h":
			out.XKpsdkH = string(in.String())
		case "x-kpsdk-fc":
			out.XKpsdkFc = string(in.String())
		case "x-kpsdk-im":
			out.XKpsdkIm = string(in.String())
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjsonD2b7633eEncodeGithubComHyperSolutionsHyperSdkGo7(out *jwriter.Writer, in KasadaHeaders) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"x-kpsdk-ct\":"
		out.RawString(prefix[1:])
		out.String(string(in.XKpsdkCt))
	}
	{
		const prefix string = ",\"x-kpsdk-dt\":"
		out.RawString(prefix)
		out.String(string(in.XKpsdkDt))
	}
	{
		const prefix string = ",\"x-kpsdk-v\":"
		out.RawString(prefix)
		out.String(string(in.XKpsdkV))
	}
	{
		const prefix string = ",\"x-kpsdk-r\":"
		out.RawString(prefix)
		out.String(string(in.XKpsdkR))
	}
	{
		const prefix string = ",\"x-kpsdk-dv\":"
		out.RawString(prefix)
		out.String(string(in.XKpsdkDv))
	}
	{
		const prefix string = ",\"x-kpsdk-h\":"
		out.RawString(prefix)
		out.String(string(in.XKpsdkH))
	}
	{
		const prefix string = ",\"x-kpsdk-fc\":"
		out.RawString(prefix)
		out.String(string(in.XKpsdkFc))
	}
	{
		const prefix string = ",\"x-kpsdk-im\":"
		out.RawString(prefix)
		out.String(string(in.XKpsdkIm))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v KasadaHeaders) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonD2b7633eEncodeGithubComHyperSolutionsHyperSdkGo7(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v KasadaHeaders) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonD2b7633eEncodeGithubComHyperSolutionsHyperSdkGo7(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *KasadaHeaders) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonD2b7633eDecodeGithubComHyperSolutionsHyperSdkGo7(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *KasadaHeaders) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonD2b7633eDecodeGithubComHyperSolutionsHyperSdkGo7(l, v)
}
