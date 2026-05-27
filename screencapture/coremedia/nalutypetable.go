package coremedia

//This code is just for printing human readable details about h264 nalus.
//check out these resources if you want to know more about what nalus are:
//https://yumichan.net/video-processing/video-compression/introduction-to-h264-nal-unit/
//https://www.semanticscholar.org/paper/Multiplexing-the-elementary-streams-of-H.264-video-Siddaraju-Rao/c7b0e625198b663be9d61c3ec7e1ec341627168c/figure/0

// Table Returns a table containing all h264 nalu types
func Table() []string { _ = "STUB: not implemented"; return nil }

var naluTypes = Table()

// GetNaluDetails creates a string containing length and type of a h264-nalu.
func GetNaluDetails(nalu []byte) string { _ = "STUB: not implemented"; return "" }

func getType(anInt byte) string { _ = "STUB: not implemented"; return "" }
