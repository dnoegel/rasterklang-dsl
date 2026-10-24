// DSL v2's own grammar version, distinct from package.json's version.
// Bump this whenever a change adds, removes, or changes the meaning of
// grammar-visible syntax (a new keyword, a new field, a changed precedence
// rule, ...), and add an entry to docs/dsl-changelog.md in the same change.
//
// Anything that stores DSL text (generated corpora, saved songs) can record
// this so it stays readable as the language evolves.
export const DSL_VERSION = '4.2.0';
