<script>
    import Home from "../company/Home.svelte"

    let token = localStorage.getItem("token");
    let record = "";
    let totalrecord = 100;
    let sData = "";
</script>

<Home {totalrecord} />