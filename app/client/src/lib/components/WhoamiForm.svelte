<script>
	import { createForm } from 'svelte-forms-lib';
	import Card from './Card.svelte';
	import Button from './Button.svelte';
	import * as yup from 'yup';
	import { startWhoamiChallenge } from '../whoami/whoami.gq';

	const { form, errors, handleChange, handleSubmit } = createForm({
		initialValues: {
			email: '',
		},
		validationSchema: yup.object().shape({
			email: yup.string().email().required(),
		}),
		onSubmit: async (values) => {
			try {
				const res = await startWhoamiChallenge({ variables: values });
				console.log({ res });
			} catch (error) {
				alert(error.message);
			}
		},
	});
</script>

<Card class="whoami-card">
	<h3>To get started, verify your email.</h3>
	<form on:submit={handleSubmit}>
		<label for="email">Email</label>
		<input
			id="email"
			name="email"
			on:change={handleChange}
			on:blur={handleChange}
			bind:value={$form.email}
		/>
		{#if $errors.email}
			<small>{$errors.email}</small>
		{/if}

		<Button type="submit">Continue</Button>
	</form>
</Card>

<style>
	h3 {
		margin-bottom: var(--lg);
	}

	:global(.whoami-card) {
		max-width: 300px;
	}
</style>
