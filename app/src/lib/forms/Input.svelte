<script lang="ts" generics="T">
    import type {HTMLInputTypeAttribute} from "svelte/elements";

    interface Props {
        label?: string
        type: HTMLInputTypeAttribute | null | undefined
        value: T
    }

    let {label, value = $bindable(), ...rest}: Props = $props()
    let fields = {...rest}
</script>

<label>
    {label}
    <input class="appearance-none border border-gray-300 outline-none p-2 rounded-sm w-full"
           {...fields}
           bind:value={value}>
</label>