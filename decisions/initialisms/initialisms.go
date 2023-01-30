package initialisms

//Words in names that are initialisms or acronyms (e.g., URL and NATO) should have the same case. URL should appear
//as URL or url (as in urlPony, or URLPony), never as Url. This also applies to ID when it is short for “identifier”;
//write appID instead of appId.
// * In names with multiple initialisms (e.g. XMLAPI because it contains XML and API), each letter within a given
//initialism should have the same case, but each initialism in the name does not need to have the same case.
// * In names with an initialism containing a lowercase letter (e.g. DDoS, iOS, gRPC), the initialism should appear
//as it would in standard prose, unless you need to change the first letter for the sake of exportedness. In these
//cases, the entire initialism should be the same case (e.g. ddos, IOS, GRPC).

const (
	XMLAPI = "XML API"
	xmlAPI = "XML API"
	IOS    = "iOS"
	iOS    = "iOS"
	GRPC   = "gRPC"
	gRPC   = "gRPC"
)

// Bad:
//XmlApi, XMLApi, XmlAPI, XMLapi
//xmlapi, xmlApi
//Ios, IoS
//ios
//Grpc
//grpc
