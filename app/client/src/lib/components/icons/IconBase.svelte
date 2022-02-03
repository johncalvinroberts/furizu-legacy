<script lang="ts">
	export let title = null;
	export let viewBox;
	export let fill = 'currentColor';
	export let stroke = 'currentColor';
	export let width: string;
	export let height: string;
</script>

<svg xmlns="http://www.w3.org/2000/svg" {viewBox} {fill} {stroke} {width} {height}>
	{#if title}
		<title>{title}</title>
	{/if}
	<slot />
</svg>

<style>
	svg {
		stroke-width: 2;
		width: 100%;
		height: auto;
		max-height: 100%;
	}
</style>
