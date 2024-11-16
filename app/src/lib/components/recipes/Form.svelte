<script lang="ts">
    import Input from "../../forms/Input.svelte";

    interface Recipe {
        name: string
        servings: number | null
        time: number | null
    }

    const recipe: Recipe = $state({
        name: 'wda',
        servings: null,
        time: null,
    })

    const saveRecipe = (e: SubmitEvent) => {
        e.preventDefault()
        console.log(recipe)
    }
</script>

<form onsubmit={saveRecipe}>
    <Input type="text" label="Name" name="name" bind:value={recipe.name}/>
    <Input type="number" label="Servings" name="servings" bind:value={recipe.servings}/>
    <Input type="number" label="Time" name="minutes" bind:value={recipe.time}/>
    <button type="submit">
        Create
    </button>
</form>

<pre>
{recipe.name}
</pre>