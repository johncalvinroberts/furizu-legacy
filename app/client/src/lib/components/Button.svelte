<script lang="ts">
	export let type = 'button';
</script>

<button {type}>
	<slot />
</button>

<style>
	button {
		color: var(--text);
		background-color: var(--muted);
		border: none;
		text-transform: uppercase;
		outline: none;
		display: block;
		min-width: 120px;
		transition: all 150ms ease;
		cursor: pointer;
		padding: var(--sm);
		margin-bottom: var(--sm);
		border-radius: var(--xxs);
	}

	button:disabled {
		opacity: 0.8;
	}

	button:focus:not(:disabled) {
		box-shadow: 0 0 0 4px var(--muted);
	}

	button:hover:not(:disabled) {
		box-shadow: var(--elevation-one);
	}
</style>
