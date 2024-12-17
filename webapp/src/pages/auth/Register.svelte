<script lang="ts">
    import Input from "../../lib/components/forms/Input.svelte";
    import Layout from "../../Layout.svelte";
    import SubmitIcon from "../../lib/icons/Submit.svelte";
    import Navbar from "../../lib/components/navigation/Navbar.svelte";
    import t from "../../lib/i18n/i18n.svelte.js"
    import {createRouter} from "../../lib/router.svelte.js";
    import {createUser} from "../../lib/auth/user.svelte.js";
    import food from "../../assets/images/auth/register.jpg"
    import Form from "../../lib/components/auth/Form.svelte";

    const router = createRouter()
    const user = createUser()

    let credentials = $state({
        displayName: "",
        email: "",
        password: "",
    })

    const handleSubmit = async (e: SubmitEvent) => {
        e.preventDefault()
        try {
            await user.login(credentials)
            router.redirectToNext()
        } catch (e) {
            // TODO: Show error toast
        }
    }
</script>

<Layout Header={Navbar}>
    <div class="flex flex-col flex-wrap md:h-full items-center justify-center">
        <Form imageSrc={food}
              submitIcon={SubmitIcon}
              submitLabel={t("register.actions.submit")}
              onSubmit={handleSubmit}
              title={t("register.title")}
              subtitle={t("register.subtitle", {url: '/login'})}
        >
            <div class="gap-3 grid grid-cols-1 md:grid-cols-2">
                <Input labelClass="col-span-2" label={t("register.labels.displayName")} type="text"
                       bind:value={credentials.displayName} required={true}/>
                <Input labelClass="col-span-2" label={t("register.labels.email")} type="email"
                       bind:value={credentials.email}
                       required={true}/>
                <Input labelClass="col-span-2 md:col-span-1" label={t("register.labels.password")} type="password"
                       bind:value={credentials.password} required={true}/>
                <Input labelClass="col-span-2 md:col-span-1" label={t("register.labels.confirmPassword")} type="password"
                       bind:value={credentials.password} required={true}/>
            </div>
        </Form>
    </div>
</Layout>