<script lang="ts">
	import { flip } from 'svelte/animate';
	import { fly } from 'svelte/transition';
	import { notifications } from '../stores/notifications';
</script>

<div class="notifications">
	{#each $notifications as notification (notification.id)}
		<div animate:flip class="toast {notification.type}" transition:fly={{ y: 30 }}>
			<div class="content">{notification.message}</div>
			{#if notification.icon}<i class={notification.icon} />{/if}
		</div>
	{/each}
</div>

<style>
	.notifications {
		position: fixed;
		bottom: 0;
		left: 0;
		right: 0;
		margin: 0 auto;
		padding: 0;
		z-index: 9999;
		display: flex;
		width: 100%;
		flex-direction: column;
		justify-content: flex-start;
		align-items: center;
		pointer-events: none;
		box-shadow: var(--elevation-one);
	}

	.toast {
		flex: 0 0 auto;
		margin-bottom: var(--xs);
		width: 100%;
	}

	.content {
		padding: 10px;
		display: block;
		color: white;
		font-weight: 500;
	}
	.danger {
		background-color: var(--error);
	}
	.success {
		background-color: var(--success);
	}
	.warning {
		background-color: var(--warning);
	}
	.info {
		background-color: var(--muted);
	}
	.default {
		background-color: var(--muted);
	}
</style>
