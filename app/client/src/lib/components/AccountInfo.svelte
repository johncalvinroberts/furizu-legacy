<script lang="ts">
	import { getContext } from 'svelte';
	import Card from './Card.svelte';
	import Button from './Button.svelte';
	import { logout, whoami } from '../stores/whoami';
	import KeyValue from './KeyValue.svelte';

	const { close } = getContext('simple-modal');
</script>

<Card>
	<h3>Account Information</h3>
	<KeyValue key="Email" value={$whoami.email} />
	<KeyValue key="User ID" value={$whoami.userId} />
	<div class="buttons">
		<Button on:click={() => close()}>Okay</Button>
		<button class="logout" on:click={() => logout()}> log out </button>
	</div>
</Card>

<style>
	h3 {
		margin-bottom: var(--sm);
	}

	.buttons {
		padding-top: var(--lg);
		display: flex;
		align-items: center;
		justify-content: space-between;
	}
	.logout {
		display: flex;
		align-items: center;
		padding: var(--xxs);
		height: var(--md);
	}
</style>
