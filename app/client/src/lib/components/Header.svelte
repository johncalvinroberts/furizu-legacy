<script>
	import { getContext } from 'svelte';
	import { page } from '$app/stores';
	import WhoamiForm from '$lib/components/WhoamiForm.svelte';
	const { open } = getContext('simple-modal');

	const showPopup = () => {
		open(WhoamiForm);
	};
</script>

<header class={$$props.class}>
	<div class="furizu-header">
		<a href="/"> Furizu ❅❆ </a>
	</div>
	<nav>
		<ul>
			<li class:active={$page.url.pathname === '/what-is-this'}>
				<a sveltekit:prefetch class="nav-link" href="/what-is-this">What is this?</a>
			</li>
			<li class:active={$page.url.pathname === '/whoami'}>
				<button class="nav-link" on:click={showPopup}>Login</button>
			</li>
		</ul>
	</nav>
</header>

<style>
	header {
		display: flex;
		justify-content: space-between;
		align-items: center;
		padding: var(--sm);
	}

	.furizu-header {
		font-size: var(--xl);
	}

	.furizu-header a:hover {
		color: var(--muted);
		transition: all 0.2s linear;
	}

	ul {
		position: relative;
		padding: 0;
		margin: 0;
		display: flex;
		justify-content: center;
		align-items: center;
		list-style: none;
		background-size: contain;
	}

	li {
		position: relative;
		height: 100%;
	}

	.nav-link {
		display: flex;
		align-items: center;
		padding: 0 1em;
		color: var(--text);
		font-weight: 700;
		font-size: 0.8rem;
		text-transform: uppercase;
		letter-spacing: 0.1em;
		text-decoration: none;
		transition: color 0.2s linear;
	}
	.nav-link:hover {
		color: var(--muted);
	}
</style>
