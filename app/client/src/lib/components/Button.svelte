<script lang="ts">
	import LoadingFlakes from './LoadingFlakes.svelte';
	import { createEventDispatcher } from 'svelte';
	export let type: 'button' | 'submit' | 'reset' = 'button';
	export let disabled = false;
	export let isLoading = false;

	const dispatch = createEventDispatcher();
</script>

<button {type} {disabled} on:click={(e) => dispatch('click', e)}>
	<slot />
	{#if isLoading}
		<LoadingFlakes />
	{/if}
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
		box-shadow: var(--elevation-one);
	}

	button:hover:not(:disabled) {
		box-shadow: var(--elevation-one);
	}
</style>
