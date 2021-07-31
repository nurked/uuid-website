package main

var tdta map[string]map[string]string

func initTranslation() {
	tdta = make(map[string]map[string]string)
	tdta["en-US"] = make(map[string]string)
	tdta["en-US"]["hello-uuid7"] = "Hello UUIDv7"
	tdta["en-US"]["description-uuid7"] = `A common case for modern applications is to create a unique
	identifier for use as a primary key in a database table.  This
	identifier usually implements an embedded timestamp that is sortable
	using the monotonic creation time in the most significant bits.  In
	addition the identifier is highly collision resistant, difficult to
	guess, and provides minimal security attack surfaces.  None of the
	existing UUID versions, including UUIDv1, fulfill each of these
	requirements in the most efficient possible way.`
	tdta["en-US"]["intro-uuid7"] = "Instead we are introducing UUIDv7. Take a look, we have a freshly generated UUIDv7 right here:"
	tdta["en-US"]["get-new"] = "Get new UUID"
	tdta["en-US"]["unix-timestamp"] = "Unix timestamp"
	tdta["en-US"]["precision"] = "Sub-second presicion"
	tdta["en-US"]["counter"] = "Counter"
	tdta["en-US"]["node"] = "Node"
	tdta["en-US"]["random"] = "Random"
	tdta["en-US"]["version"] = "Version"
	tdta["en-US"]["variant"] = "Variant"
	tdta["en-US"]["unix-timestamp-description"] = "no more flying DeLorians. You always in control and you know where and when this UUID happened. First 36 bits are dedicated for the Unix timestamp. No random bit ordering. No little-endian bits. It is just a number that can be parsed by any computer system."
	tdta["en-US"]["precision-description"] = "Need to be more specific about your time? UUIDv7 allows you to store up to 48 bits of sub-second precision. Not just that. The data is encoded with a trickle of magic that allows it to be binary sortable. But we are a very good magicians. We explain all our tricks."
	tdta["en-US"]["counter-description"] = "Is your flux capacitor overcharged and you are producing millions of UUIDs per milli-second? Use counter to keep them well ordered."
	tdta["en-US"]["node-description"] = "You can store an information about a machine that produced this UUID in this block."
	tdta["en-US"]["random-description"] = "Add some randomity at the end to ensure that your UUID is totally unique and no other UUIDs are just like it."
	tdta["en-US"]["version-description"] = "There are times to be boring for a reason of a backward-compatibility. This field contains the version number to be compatible with UUIDv4."
	tdta["en-US"]["variant-description"] = "Even more boring stuff for compatibility with UUIDv4."
	tdta["en-US"]["encoding-intro"] = "Special mention to our amazing counter precision encoding algorithm."
	tdta["en-US"]["footer-ietf"] = "IETF: New UUID formats"
	tdta["en-US"]["footer-github-draft"] = "Github: uuid6-ietf-draft"
	tdta["en-US"]["footer-golang-gen"] = "golang generator uuid6go-proto"
	tdta["en-US"]["footer-vugu"] = "site is written in vugu"
	tdta["en-US"]["footer-contact"] = "Contact:"
	tdta["en-US"]["footer-peabody"] = "Brad G. Peabody"
	tdta["en-US"]["footer-davis"] = "Kyzer R. Davis"
	tdta["en-US"]["footer-and"] = "and"
	tdta["en-US"]["footer-copyright"] = "Copyright"

}

func translate(lang string, str string) string {

	l, ok := tdta[lang]
	if !ok {
		return "!!!NOLANG:" + lang + "!!!"
	}

	s, ok := l[str]
	if !ok {
		return "!!!NOSTR:" + lang + "|" + str + "!!!"
	}

	return s
}
