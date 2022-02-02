<script context="module" lang="ts">
	export const prerender = true;
	export async function load({ fetch }) {
		// load here?
		return {};
	}
</script>

<div class="home-wrapper">
	<h1>Welcome to Furizu ❅❆</h1>
	<h3>Upload your stuff into resilient <strong>cloud cold storage</strong>. Forever.</h3>
</div>

<style>
	.home-wrapper {
		padding: var(--xl);
		display: flex;
		align-items: center;
		justify-content: center;
		flex-wrap: wrap;
	}
	.home-wrapper h1,
	.home-wrapper h3 {
		display: block;
		width: 100%;
		text-align: center;
	}
	h1 {
		font-weight: bold;
		font-size: 3rem;
		letter-spacing: -0.08rem;
		margin-bottom: var(--md);
	}
</style>
