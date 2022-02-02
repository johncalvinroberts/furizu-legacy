import adapter from '@sveltejs/adapter-static';
import preprocess from 'svelte-preprocess';
import gQueryCodegen from '@leveluptuts/g-query/codegen';

/** @type {import('@sveltejs/kit').Config} */
const config = {
	// Consult https://github.com/sveltejs/svelte-preprocess
	// for more information about preprocessors
	preprocess: preprocess(),

	kit: {
		adapter: adapter({
			// default options are shown
			fallback: null,
		}),
		// hydrate the <div id="svelte"> element in src/app.html
		target: '#svelte',
		vite: {
			plugins: [
				gQueryCodegen({
					schema: '../graph/schema.graphql',
					out: './src/lib/graphql',
					gPath: '$lib/g',
				}),
			],
		},
	},
};

export default config;
