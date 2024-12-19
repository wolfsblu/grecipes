<script lang="ts" generics="T">
    import type {HTMLInputTypeAttribute} from "svelte/elements";

    interface Props {
        hint?: string
        inputClass?: string
        label?: string
        labelClass?: string
        required?: boolean
        type: HTMLInputTypeAttribute | null | undefined
        value: T
    }

    let {hint, inputClass, label, labelClass, value = $bindable(), ...rest}: Props = $props()
    let fields = {...rest}
</script>

<label class={labelClass}>
    {label}
    <input class="appearance-none border border-gray-200 outline-none p-2 rounded w-full {inputClass}"
           {...fields}
           bind:value={value}>
</label>

{#if hint}
    <p class="text-xs">{@html hint}</p>
{/if}