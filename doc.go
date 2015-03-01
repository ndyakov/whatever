// Whatever
// Whatever is a package that holds a single type
// Params. The Params type is map[string]interface{} with
// some useful methods to it.
// This type was initially created to be used as structure
// in which JSON requests are to be unmarshaled and accessed
// with ease. Because of this, the Params type have a getter
// for time.Time that will parse a Date string that follows
// the RFC3339 format (the Javascript build-in JSON format).
// There is a method that can transform the Params structure to
// url.Values structure with specified prefix and suffix, for the
// result can be used with Gorilla`s schema or Goji`s params packages.
// Although some of the getters are useful for unmarshaled JSON
// date you can also Add your own values to the Params structure.
// You can also access nested Params objects.
// If you need you can validate the existence of a specific key by
// using the Required method.
package whatever
