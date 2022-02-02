<script>
	import { createForm } from 'svelte-forms-lib';
	import Card from './Card.svelte';
	import * as yup from 'yup';

	const { form, errors, handleChange, handleSubmit } = createForm({
		initialValues: {
			email: '',
		},
		validationSchema: yup.object().shape({
			email: yup.string().email().required(),
		}),
		onSubmit: (values) => {
			alert(JSON.stringify(values));
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

		<button type="submit">Continue</button>
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
