<script lang="ts">
	export let key: string = '';
	export let value: any = '';
</script>

<div class="kv">
	<div class="key">{key}</div>
	<div class="value">
		{value}
	</div>
</div>

<style>
	.kv {
		display: flex;
		justify-content: flex-start;
		align-items: center;
		font-size: var(--sm);
		color: var(--muted);
		padding: var(--xs) var(--sm);
		margin-bottom: var(--xs);
	}

	.key {
		flex: 0 0 100px;
		border-right: solid 1px var(--background);
	}
	.value {
		padding-left: var(--xs);
	}
</style>
