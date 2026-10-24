// Staleness check for the generated file docs/dsl-spec.generated.json
// (scripts/generate-dsl-spec.mjs): regenerates it in memory and diffs
// against what's actually committed. If this fails, someone changed
// src/dsl/{keywords,fieldDocs,version}.js (or the things they derive from)
// without re-running `npm run generate:dsl-spec` - the standard "committed
// generated file must match its generator" pattern, same idea as a
// lockfile-freshness check.
import { describe, expect, it } from 'vitest';
import { readFileSync } from 'node:fs';
import { fileURLToPath } from 'node:url';
import { dirname, join } from 'node:path';
import { serializeDslSpec } from '../../scripts/generate-dsl-spec.mjs';

const __dirname = dirname(fileURLToPath(import.meta.url));
const specPath = join(__dirname, '..', '..', 'docs', 'dsl-spec.generated.json');

describe('docs/dsl-spec.generated.json', () => {
  it('matches what scripts/generate-dsl-spec.mjs would produce right now', () => {
    const committed = readFileSync(specPath, 'utf8');
    expect(committed, 'docs/dsl-spec.generated.json is stale -- run `npm run generate:dsl-spec` and commit the result').toBe(
      serializeDslSpec(),
    );
  });

  it('is valid JSON matching the current DSL_VERSION', async () => {
    const { DSL_VERSION } = await import('../../src/dsl/version.js');
    const parsed = JSON.parse(readFileSync(specPath, 'utf8'));
    expect(parsed.dslVersion).toBe(DSL_VERSION);
  });
});
