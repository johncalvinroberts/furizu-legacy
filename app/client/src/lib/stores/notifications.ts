import { writable, derived } from 'svelte/store';

enum NotificationType {
	default = 'default',
	danger = 'danger',
	warning = 'warning',
	info = 'info',
	success = 'success',
}

type Notification = {
	id: string;
	message: string;
	type: NotificationType;
	timeout: number;
};

const DEFAULT_TIMEOUT = 5000;

function id() {
	return Math.random().toString(36).substring(2, 9);
}

function createNotificationStore() {
	const notificationsStore = writable<Notification[]>([]);

	function send(
		message: string,
		type: NotificationType = NotificationType.default,
		timeout: number
	) {
		notificationsStore.update((state) => {
			return [...state, { id: id(), type, message, timeout }];
		});
	}

	const notifications = derived(notificationsStore, ($notificationsStore, set) => {
		set($notificationsStore);
		if ($notificationsStore.length > 0) {
			const timer = setTimeout(() => {
				notificationsStore.update((state) => {
					state.shift();
					return state;
				});
			}, $notificationsStore[0].timeout);
			return () => {
				clearTimeout(timer);
			};
		}
	});

	const { subscribe } = notifications;

	return {
		subscribe,
		send,
		default: (msg: string, timeout: number = DEFAULT_TIMEOUT) =>
			send(msg, NotificationType.default, timeout),
		danger: (msg: string, timeout: number = DEFAULT_TIMEOUT) =>
			send(msg, NotificationType.danger, timeout),
		warning: (msg: string, timeout: number = DEFAULT_TIMEOUT) =>
			send(msg, NotificationType.warning, timeout),
		info: (msg: string, timeout: number = DEFAULT_TIMEOUT) =>
			send(msg, NotificationType.info, timeout),
		success: (msg: string, timeout: number = DEFAULT_TIMEOUT) =>
			send(msg, NotificationType.success, timeout),
	};
}

export const notifications = createNotificationStore();
