<script>
    import Layout from "../../Layout.svelte";
    import Navbar from "../../lib/components/navigation/Navbar.svelte";
    import Form from "../../lib/components/recipes/Form.svelte";
</script>

<Layout Header={Navbar}>
    <div class="border rounded border-gray-300 p-3">
        <Form />
    </div>
</Layout>
