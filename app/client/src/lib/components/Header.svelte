<script>
	import { page } from '$app/stores';
</script>

<header class={$$props.class}>
	<div class="furizu-header">
		<a href="/"> Furizu ❅❆ </a>
	</div>
	<nav>
		<ul>
			<li class:active={$page.url.pathname === '/what-is-this'}>
				<a sveltekit:prefetch href="/what-is-this">What is this?</a>
			</li>
			<li class:active={$page.url.pathname === '/whoami'}>
				<a sveltekit:prefetch href="/whoami">Login</a>
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
	.furizu-header a {
		text-decoration: none;
		transition: color 0.2s linear;
	}

	.furizu-header a:hover {
		color: var(--muted);
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

	nav a {
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

	a:hover {
		color: var(--muted);
	}
</style>
