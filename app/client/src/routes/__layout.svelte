<script lang="ts">
	import '../styles/fonts.css';
	import '../styles/global.css';
	import Header from '../lib/components/Header.svelte';
	import Modal from '../lib/components/Modal.svelte';
	import Toast from '../lib/components/Toast.svelte';

	const year = new Date().getFullYear();
</script>

<Modal>
	<Header class="header" />
	<main>
		<slot />
	</main>

	<footer>
		© furizu {year}
	</footer>
	<Toast />
</Modal>

<style>
	:global(.header) {
		position: sticky;
		top: 0;
		backdrop-filter: blur(12px);
	}

	main {
		padding: var(--md);
		width: 100%;
		max-width: 1024px;
		margin: 0 auto;
		box-sizing: border-box;
		min-height: calc(100vh - 104px);
		flex: 1;
		justify-content: center;
		display: flex;
		align-items: center;
	}

	footer {
		display: flex;
		flex-direction: column;
		justify-content: center;
		align-items: center;
		font-size: var(--xs);
		height: 40px;
	}
</style>
