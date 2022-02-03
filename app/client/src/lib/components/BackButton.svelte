<script>
	import { createEventDispatcher } from 'svelte';
	import ArrowLeft from './icons/ArrowLeft.svelte';
	const dispatch = createEventDispatcher();
</script>

<button class="back-button" on:click={(e) => dispatch('click', e)}>
	<span>
		<ArrowLeft />
	</span>
	Back
</button>

<style>
	.back-button {
		display: flex;
		align-items: center;
		padding: var(--xxs);
		height: var(--md);
	}

	.back-button:hover > span {
		margin-right: 4px;
	}

	.back-button > span {
		margin-right: 2px;
		transition: margin 0.1s linear;
	}
</style>
