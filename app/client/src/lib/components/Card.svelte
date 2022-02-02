<script lang="ts">
</script>

<div class={`${$$props.class} card`}>
	<slot />
</div>

<style>
	.card {
		background-color: white;
		border-radius: 2px;
		padding: 15px;
		min-width: 265px;
		transition: all 150ms ease;
		/* box-shadow: var(--elevation-one); */
		color: var(--background);
	}
</style>
