import { writable } from 'svelte/store';

export const themeChanged = writable(Date.now());

export function notifyThemeChange() {
    themeChanged.set(Date.now());
}