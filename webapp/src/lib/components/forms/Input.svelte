<script lang="ts" generics="T">
    import type {HTMLInputTypeAttribute} from "svelte/elements";

    interface Props {
        class?: string
        label?: string
        required?: boolean
        type: HTMLInputTypeAttribute | null | undefined
        value: T
    }

    let {class: cls, label, value = $bindable(), ...rest}: Props = $props()
    let fields = {...rest}
</script>

<label>
    {label}
    <input class="appearance-none border border-gray-200 outline-none p-2 rounded w-full {cls}"
           {...fields}
           bind:value={value}>
</label>