import "iter"

func Lines(s string) iter.Seq[string]
func SplitSeq(s, sep string) iter.Seq[string]
func SplitAfterSeq(s, sep string) iter.Seq[string]
func FieldsSeq(s string) iter.Seq[string]
func FieldsFuncSeq(s string, f func(rune) bool) iter.Seq[string]
