<script lang="ts">
	import Card from './Card.svelte';
	import Button from './Button.svelte';
	import { EMAIL_REGEX } from '../constants';
	import { startWhoamiChallenge, redeemWhoamiChallenge } from '../whoami/whoami.gq';
	import { notifications } from '$lib/stores/notifications';
	import Check from './icons/Check.svelte';
	import BackButton from './BackButton.svelte';
	import { setWhoamiState } from '../stores/whoami';

	enum steps {
		SEND_WHOAMI_CHALLENGE,
		REDEEM_WHOAMI_CHALLENGE,
		AUTHENTICATED,
	}

	// let step: steps = steps.SEND_WHOAMI_CHALLENGE;
	let step: steps = steps.AUTHENTICATED;
	let isLoading = false;
	let email = '';
	let emailError = '';
	let code = '';
	let redeemError = '';

	const handleSubmitEmail = async (e: SubmitEvent) => {
		e.preventDefault();
		try {
			if (!EMAIL_REGEX.test(email)) {
				emailError = 'Please enter a valid email.';
				return;
			} else {
				emailError = '';
			}
			isLoading = true;
			const res = await startWhoamiChallenge({ variables: { email } });
			if (res.errors?.length > 0) {
				throw new Error(res.errors[0]?.message);
			}
			notifications.success('Email sent. Please check your email.');
			step = steps.REDEEM_WHOAMI_CHALLENGE;
		} catch (error) {
			notifications.danger(`Something went wrong: ${error.message}`);
		}
		isLoading = false;
	};

	const handleRedeem = async (e: SubmitEvent) => {
		e.preventDefault();
		try {
			isLoading = true;
			const res = await redeemWhoamiChallenge({ variables: { email, token: code } });
			if (res.errors?.length > 0) {
				throw new Error(res.errors[0]?.message);
			}
			const { refreshToken, accessToken } = res.redeemWhoamiChallenge;
			setWhoamiState({ refreshToken, accessToken });
			notifications.success('Signed in.');
			step = steps.AUTHENTICATED;
		} catch (error) {
			notifications.danger(`Something went wrong: ${error.message}`);
		}
		isLoading = false;
	};
</script>

<Card class="whoami-card">
	{#if step === steps.SEND_WHOAMI_CHALLENGE}
		<h3>To get started, verify your email.</h3>
		<form on:submit={handleSubmitEmail}>
			<label for="email">Email</label>
			<input id="email" name="email" bind:value={email} />
			{#if emailError}
				<small>{emailError}</small>
			{/if}
			<Button type="submit" {isLoading}>Continue</Button>
		</form>
	{/if}
	{#if step === steps.REDEEM_WHOAMI_CHALLENGE}
		<h3>
			<span class="success-icon">
				<Check />
			</span>
			Good job. Enter the code from your email.
		</h3>
		<form on:submit={handleRedeem}>
			<label for="email">Email</label>
			<input id="email" name="email" bind:value={email} disabled />
			<label for="code">Verification Code</label>
			<input id="code" name="code" bind:value={code} />
			{#if redeemError}
				<small>{redeemError}</small>
			{/if}
			<div class="buttons">
				<Button type="submit" {isLoading}>Complete</Button>
				<BackButton on:click={() => (step = steps.SEND_WHOAMI_CHALLENGE)} />
			</div>
		</form>
	{/if}
	{#if step === steps.AUTHENTICATED}
		<div class="authenticated">
			<span class="success-icon-authenticated">
				<Check />
			</span>
			<h3>Done. You're authenticated.</h3>
		</div>
		<div class="authenticated-buttons">
			<Button>Okay.</Button>
		</div>
	{/if}
</Card>

<style>
	h3 {
		margin-bottom: var(--lg);
	}

	:global(.whoami-card) {
		max-width: 300px;
	}
	.success-icon {
		color: var(--success);
		display: inline-flex;
		justify-content: center;
		align-items: center;
	}
	.buttons {
		display: flex;
		justify-content: space-between;
	}
	.authenticated {
		display: flex;
		justify-content: space-between;
	}
	.success-icon-authenticated {
		width: 100px;
		margin-right: var(--md);
		color: var(--success);
	}

	.authenticated-buttons {
		justify-content: center;
		display: flex;
	}
</style>
