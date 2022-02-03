<span class="loading">
	<span class="flake"> ❆ </span>
	<span class="flake"> ❅ </span>
</span>

<style>
	.loading {
		margin-left: var(--xxs);
		position: relative;
		width: 80px;
		height: 80px;
	}

	.loading .flake {
		position: absolute;
		animation-timing-function: cubic-bezier(0, 1, 1, 0);
	}

	.loading .flake:nth-child(1) {
		animation: shuffle1 0.8s infinite;
		left: 0;
	}

	.loading .flake:nth-child(2) {
		animation: shuffle2 0.8s infinite;
	}

	@keyframes shuffle1 {
		0% {
			transform: scale(0);
		}
		100% {
			transform: scale(1);
		}
	}
	@keyframes shuffle2 {
		0% {
			transform: scale(1);
		}
		100% {
			transform: scale(0);
		}
	}
</style>
