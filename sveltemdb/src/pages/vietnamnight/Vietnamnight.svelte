<script>
    import Home from "../vietnamnight/Home.svelte";
    
    export let table_header_font = "";
	export let table_body_font = "";
    
    let token = localStorage.getItem("token");
    let akses_page = true;
    let listHome = [];
    let sData = "";
    let record = "";
    let record_message = "";
    let totalrecord = 0;

    async function initapp() {
        const res = await fetch("/api/valid", {
            method: "POST",
            headers: {
                "Content-Type": "application/json",
                Authorization: "Bearer " + token,
            },
            body: JSON.stringify({
                page: "VIETNAMNIGHT-VIEW",
            }),
        });
        const json = await res.json();
        if (json.status === 400) {
            logout();
        } else if (json.status == 403) {
            alert(json.message);
            akses_page = false;
        } else {
            initHome();
        }
    }
    async function initHome() {
        const res = await fetch("/api/vietnamnight", {
            method: "POST",
            headers: {
                "Content-Type": "application/json",
                Authorization: "Bearer " + token,
            },
            body: JSON.stringify({
            }),
        });
        const json = await res.json();
        if (json.status == 200) {
            record = json.record;
            record_message = json.message;
            if (record != null) {
                totalrecord = record.length;
                let no = 0
                for (var i = 0; i < record.length; i++) {
                    no = no + 1;
                    listHome = [
                        ...listHome,
                        {
                            vietnamnight_no: no,
                            vietnamnight_id: record[i]["vietnamnight_id"],
                            vietnamnight_date: record[i]["vietnamnight_date"],
                            vietnamnight_prize1_1300: record[i]["vietnamnight_prize1_1300"],
                            vietnamnight_prize1_1700: record[i]["vietnamnight_prize1_1700"],
                            vietnamnight_prize1_2000: record[i]["vietnamnight_prize1_2000"],
                            vietnamnight_prize1_2200: record[i]["vietnamnight_prize1_2200"],
                            vietnamnight_create: record[i]["vietnamnight_create"],
                            vietnamnight_update: record[i]["vietnamnight_update"],
                        },
                    ];
                }
            }
        } else {
            logout();
        }
    }
    async function logout() {
        localStorage.clear();
        window.location.href = "/";
    }
    const handleEditData = (e) => {
        admin_username = e.detail.e;
        sData = "Edit";
        alert(admin_username)
        // editAdmin(admin_username);
    };
    const handleRefreshData = (e) => {
        listHome = [];
        totalrecord = 0;
        setTimeout(function () {
            initHome();
        }, 500);
    };
    initapp()
</script>

<Home
    on:handleEditData={handleEditData}
    on:handleRefreshData={handleRefreshData}
    {token}
    {table_header_font}
    {table_body_font}
    {listHome}
    {totalrecord}
/>