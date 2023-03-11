<script>
    import { Input } from "sveltestrap";
    import { initializeApp } from "firebase/app";
    import { getDatabase, ref, set } from "firebase/database";
    import dayjs from "dayjs";
    
    import Panel from "../../components/Panel.svelte";
    import Loader from "../../components/Loader.svelte";
	import Button from "../../components/Button.svelte";
	import Modal from "../../components/Modal.svelte";
    import { createEventDispatcher } from "svelte";

    const firebaseConfig = {
        apiKey: "AIzaSyBqVRbGvJBb1JEfKYyN6jgocZjzsx2lN2A",
        authDomain: "dazzling-pillar-328210.firebaseapp.com",
        databaseURL: "https://dazzling-pillar-328210-default-rtdb.asia-southeast1.firebasedatabase.app",
        projectId: "dazzling-pillar-328210",
        storageBucket: "dazzling-pillar-328210.appspot.com",
        messagingSenderId: "770359422621",
        appId: "1:770359422621:web:7933922e00547dc735ee74"
    };
    const app = initializeApp(firebaseConfig);
    const db = getDatabase(app);
	export let table_header_font
	export let table_body_font
	export let token
	export let listHome = []
	export let totalrecord = 0
    let dispatch = createEventDispatcher();
	let title_page = "VIETNAM - NIGHT"
    let sData = "";
    let myModal_newentry = "";
    let tanggal_keluaran = "";
    let date_keluaran ="";
    let idrecord = 0;
    let prize1_1300 = "";
    let prize1_1700 = "";
    let prize1_2000 = "";
    let prize1_2200 = "";
    let prize1_1300_flag = false;
    let prize1_1700_flag = false;
    let prize1_2000_flag = false;
    let prize1_2200_flag = false;
    let prize1_1300_save_flag = false;
    let prize1_1700_save_flag = false;
    let prize1_2000_save_flag = false;
    let prize1_2200_save_flag = false;
    let css_loader = "display: none;";
    let msgloader = "";

    
    const NewData = () => {
        clearField()
        sData = "New"
        myModal_newentry = new bootstrap.Modal(document.getElementById("modalentry"));
        myModal_newentry.show();
        
    };
    const RefreshHalaman = () => {
        dispatch("handleRefreshData", "call");
    };
    const EditData = (e,tanggal,prize_1,prize_2,prize_3,prize_4) => {
        sData = "Edit"
        idrecord = parseInt(e)
        prize1_1300 = prize_1
        prize1_1700 = prize_2
        prize1_2000 = prize_3
        prize1_2200 = prize_4
        if(prize1_1300 !=""){
            prize1_1300_flag = true;
            prize1_1300_save_flag = true;
        }else{
            prize1_1300_flag = false;
            prize1_1300_save_flag = false;
        }
        if(prize1_1700 !=""){
            prize1_1700_flag = true;
            prize1_1700_save_flag = true;
        }else{
            prize1_1700_flag = false;
            prize1_1700_save_flag = false;
        }
        if(prize1_2000 !=""){
            prize1_2000_flag = true;
            prize1_2000_save_flag = true;
        }else{
            prize1_2000_flag = false;
            prize1_2000_save_flag = false;
        }
        if(prize1_2200 !=""){
            prize1_2200_flag = true;
            prize1_2200_save_flag = true;
        }else{
            prize1_2200_flag = false;
            prize1_2200_save_flag = false;
        }
        tanggal_keluaran = tanggal;
        myModal_newentry = new bootstrap.Modal(document.getElementById("modalentryedit"));
        myModal_newentry.show();
    };
    async function handleSave() {
        let flag = true
        let msg = ""
        if(date_keluaran == ""){
            flag = false
            msg = "The Date is required"
        }
        if(flag){
            const res = await fetch("/api/savevietnamnight", {
                method: "POST",
                headers: {
                    "Content-Type": "application/json",
                    Authorization: "Bearer " + token,
                },
                body: JSON.stringify({
                    sdata: sData,
                    page:"SDSB4DNIGHT-SAVE",
                    idrecord: parseInt(0),
                    tanggal: date_keluaran,
                }),
            });
            const json = await res.json();
            if (json.status == 200) {
                set(ref(db, 'vietnamnight'), {
                    datedraw: dayjs(date_keluaran).format("DD-MMM-YYYY"),
                    nextdraw: dayjs(date_keluaran).add(1,'day').format("YYYY-MM-DD"),
                    prize1_1300: "",
                    prize1_1700: "",
                    prize1_2000: "",
                    prize1_2200: "",
                });

                msgloader = json.message;
                myModal_newentry.hide()
                RefreshHalaman()
            } else if(json.status == 403){
                alert(json.message)
            } else {
                msgloader = json.message;
            }
            setTimeout(function () {
                css_loader = "display: none;";
            }, 1000);
        }else{
            alert(msg)
        }
    }
    async function handleSaveGenerator(tipe,prize) {
        let flag_save = false;
        if(prize.length == 6){
            flag_save = true;
        }
        if(flag_save){
            const res = await fetch("/api/savegeneratorvietnamnight", {
                method: "POST",
                headers: {
                    "Content-Type": "application/json",
                    Authorization: "Bearer " + token,
                },
                body: JSON.stringify({
                    sdata: sData,
                    page:"SDSB4DNIGHT-SAVE",
                    idrecord: parseInt(idrecord),
                    tipe: tipe,
                    prize: prize.toString(),
                }),
            });
            const json = await res.json();
            if (json.status == 200) {
                msgloader = json.message;
                if(msgloader == "Success"){
                    RefreshHalaman()
                    switch(tipe){
                        case "prize1_1300":
                            prize1_1300_save_flag = true;
                            prize1_1300_flag = true;
                            
                            break;
                        case "prize1_1700":
                            prize1_1700_save_flag = true;
                            prize1_1700_flag = true;
                            break;
                        case "prize1_2000":
                            prize1_2000_save_flag = true;
                            prize1_2000_flag = true;
                            break;
                        case "prize1_2200":
                            prize1_2200_save_flag = true;
                            prize1_2200_flag = true;
                            break;
                    }
                    set(ref(db, 'vietnamnight'), {
                        datedraw: dayjs(tanggal_keluaran).format("DD-MMM-YYYY"),
                        nextdraw: dayjs(tanggal_keluaran).add(1,'day').format("YYYY-MM-DD"),
                        prize1_1300: prize1_1300,
                        prize1_1700: prize1_1700,
                        prize1_2000: prize1_2000,
                        prize1_2200: prize1_2200,
                    });
                }else{
                    alert(msgloader)
                }
                
            } else if(json.status == 403){
                alert(json.message)
            } else {
                msgloader = json.message;
            }
            setTimeout(function () {
                css_loader = "display: none;";
            }, 1000);
        }else{
            alert("Minimal 6 Digit")
        }
        
    }
    async function handleGeneratorAutomation() {
        const res = await fetch("/api/generatornumbervietnamnight", {
            method: "POST",
            headers: {
                "Content-Type": "application/json",
                Authorization: "Bearer " + token,
            },
        });
        const json = await res.json();
        if (json.status == 200) {
            msgloader = json.message;
            RefreshHalaman()
        } else if(json.status == 403){
            alert(json.message)
        } else {
            msgloader = json.message;
        }
        setTimeout(function () {
            css_loader = "display: none;";
        }, 1000);
    }
    function clearField(){
        date_keluaran = ""
    }
    function generate(field){
        let numbergenerate = (Math.floor(Math.random() * 1000000) + 1000000).toString().substring(1);
        switch(field){
            case "prize1_1300":
                prize1_1300 = numbergenerate
                break;
            case "prize1_1700":
                prize1_1700 = numbergenerate
                break;
            case "prize1_2000":
                prize1_2000 = numbergenerate
                break;
            case "prize1_2200":
                prize1_2200 = numbergenerate
                break;
        }
    }
    function callFunction(event){
        switch(event.detail){
            case "NEW":
                NewData();
                break;
            case "REFRESH":
                RefreshHalaman();break;
            case "GENERATOR":
                handleGeneratorAutomation();break;
            case "SAVE":
                handleSubmit();break;
        }
    }
    const handleKeyboard_format = () => {
		let numbera;
		for (let i = 0; i < prize1_1300.length; i++) {
			numbera = parseInt(prize1_1300[i]);
			if (isNaN(numbera)) {
				prize1_1300 = "";
			}
		}
        for (let i = 0; i < prize1_1700.length; i++) {
			numbera = parseInt(prize1_1700[i]);
			if (isNaN(numbera)) {
				prize1_1700 = "";
			}
		}
        for (let i = 0; i < prize1_2000.length; i++) {
			numbera = parseInt(prize1_2000[i]);
			if (isNaN(numbera)) {
				prize1_2000 = "";
			}
		}
        for (let i = 0; i < prize1_2200.length; i++) {
			numbera = parseInt(prize1_2200[i]);
			if (isNaN(numbera)) {
				prize1_2200 = "";
			}
		}
    }
</script>
<div id="loader" style="margin-left:50%;{css_loader}">
    {msgloader}
</div>
<div class="container" style="margin-top: 70px;">
    <div class="row">
        <div class="col-sm-12">
            <Button
                on:click={callFunction}
                button_function="NEW"
                button_title="New"
                button_css="btn-dark"/>
            <Button
                on:click={callFunction}
                button_function="REFRESH"
                button_title="Refresh"
                button_css="btn-primary"/>
            <Button
                on:click={callFunction}
                button_function="GENERATOR"
                button_title="Generator"
                button_css="btn-primary"/>
            <Panel
                card_title="{title_page}"
                card_footer={totalrecord}>
                <slot:template slot="card-body">
                        <table class="table table-striped table-hover">
                            <thead>
                                <tr>
                                    <th NOWRAP width="1%" style="text-align: center;vertical-align: top;">&nbsp;</th>
                                    <th NOWRAP width="1%" style="text-align: center;vertical-align: top;font-weight:bold;font-size:{table_header_font};">NO</th>
                                    <th NOWRAP width="*" style="text-align: center;vertical-align: top;font-weight:bold;font-size: {table_header_font};">DATE</th>
                                    <th NOWRAP width="20%" style="text-align: center;vertical-align: top;font-weight:bold;font-size: {table_header_font};">PRIZE 1 - 13:00</th>
                                    <th NOWRAP width="20%" style="text-align: center;vertical-align: top;font-weight:bold;font-size: {table_header_font};">PRIZE 1 - 17:00</th>
                                    <th NOWRAP width="20%" style="text-align: center;vertical-align: top;font-weight:bold;font-size: {table_header_font};">PRIZE 1 - 20:00</th>
                                    <th NOWRAP width="20%" style="text-align: center;vertical-align: top;font-weight:bold;font-size: {table_header_font};">PRIZE 1 - 22:00</th>
                                    <th NOWRAP width="10%" style="text-align: left;vertical-align: top;font-weight:bold;font-size: {table_header_font};">CREATE</th>
                                    <th NOWRAP width="10%" style="text-align: left;vertical-align: top;font-weight:bold;font-size: {table_header_font};">UPDATE</th>
                                </tr>
                            </thead>
                            {#if totalrecord > 0}
                            <tbody>
                                {#each listHome as rec }
                                    <tr>
                                        <td NOWRAP style="text-align: center;vertical-align: top;cursor:pointer;">
                                            <i 
                                                on:click={() => {
                                                    EditData(rec.vietnamnight_id,rec.vietnamnight_date,
                                                    rec.vietnamnight_prize1_1300,
                                                    rec.vietnamnight_prize1_1700,
                                                    rec.vietnamnight_prize1_2000,
                                                    rec.vietnamnight_prize1_2200);
                                                }} 
                                                class="bi bi-pencil"></i>
                                        </td>
                                        <td NOWRAP style="text-align: center;vertical-align: top;font-size: {table_body_font};">{rec.vietnamnight_no} - {rec.vietnamnight_id}</td>
                                        <td NOWRAP style="text-align: center;vertical-align: top;font-size: {table_body_font};">{rec.vietnamnight_date}</td>
                                        <td NOWRAP style="text-align: center;vertical-align: top;font-size: {table_body_font};">{rec.vietnamnight_prize1_1300}</td>
                                        <td NOWRAP style="text-align: center;vertical-align: top;font-size: {table_body_font};">{rec.vietnamnight_prize1_1700}</td>
                                        <td NOWRAP style="text-align: center;vertical-align: top;font-size: {table_body_font};">{rec.vietnamnight_prize1_2000}</td>
                                        <td NOWRAP style="text-align: center;vertical-align: top;font-size: {table_body_font};">{rec.vietnamnight_prize1_2200}</td>
                                        <td NOWRAP style="text-align: left;vertical-align: top;font-size: {table_body_font};">{rec.vietnamnight_create}</td>
                                        <td NOWRAP style="text-align: left;vertical-align: top;font-size: {table_body_font};">{rec.vietnamnight_update}</td>
                                    </tr>
                                {/each}
                            </tbody>
                            {:else}
                            <tbody>
                                <tr>
                                    <td colspan="10">
                                        <center>
                                            <Loader />
                                        </center>
                                    </td>
                                </tr>
                            </tbody>
                            {/if} 
                        </table>
                </slot:template>
            </Panel>
        </div>
    </div>
</div>

<Modal
	modal_id="modalentry"
	modal_size="modal-dialog-centered"
	modal_title="{title_page+"/"+sData}"
    modal_footer_css="padding:5px;"
	modal_footer={true}>
	<slot:template slot="body">
        <div class="mb-3">
            <label for="exampleForm" class="form-label">Date</label>
			<Input
                bind:value={date_keluaran}
                type="date"
                name="date"
                id="exampleDate"
                data-date-format="dd-mm-yyyy"
                placeholder="date placeholder"/>
		</div>
	</slot:template>
	<slot:template slot="footer">
        <Button
            on:click={() => {
                handleSave();
            }} 
            button_function="SAVE"
            button_title="Save"
            button_css="btn-warning"/>
	</slot:template>
</Modal>

<Modal
	modal_id="modalentryedit"
	modal_size="modal-dialog-centered"
	modal_title="{title_page+"/"+sData}"
    modal_footer_css="padding:5px;"
	modal_footer={false}>
	<slot:template slot="body">
        <div class="mb-3">
            <label for="exampleForm" class="form-label">Date</label>
			<Input
                bind:value={tanggal_keluaran}
                type="text"
                name="date"
                id="exampleDate"
                disabled
                placeholder="Tanggal"/>
		</div>
        
        <div class="mb-3">
            <label for="exampleForm" class="form-label">
                <div class="d-flex flex-row">
                    <div class="">Prize 13:00</div>
                    <div class="ps-2 mt-3" style="font-size: 11px;color:blue;">OPEN: 12:55 - 13:00</div>
                </div>
            </label>
            <div class="input-group mb-3">
                <Input
                    bind:value={prize1_1300}
                    on:keyup={handleKeyboard_format}    
                    disabled='{prize1_1300_flag}'
                    type="text"
                    minlength="6"
                    maxlength="6"
                    placeholder="Prize 13:00"/>
                <button
                    on:click={() => {
                        generate("prize1_1300");
                    }}  
                    disabled='{prize1_1300_save_flag}' 
                    type="button" class="btn btn-info">Generate</button>
                <button
                    on:click={() => {
                        handleSaveGenerator("prize1_1300",prize1_1300);
                    }} 
                    disabled='{prize1_1300_save_flag}' 
                    type="button" class="btn btn-warning">Save</button>
            </div>
		</div>
        <div class="mb-3">
            <label for="exampleForm" class="form-label">
                <div class="d-flex flex-row">
                    <div class="">Prize 17:00</div>
                    <div class="ps-2 mt-3" style="font-size: 11px;color:blue;">OPEN: 16:55 - 17:00</div>
                </div>
            </label>
            <div class="input-group mb-3">
                <Input
                    bind:value={prize1_1700}
                    on:keyup={handleKeyboard_format}    
                    disabled='{prize1_1700_flag}'
                    type="text"
                    minlength="6"
                    maxlength="6"
                    placeholder="Prize 17:00"/>
                <button
                    on:click={() => {
                        generate("prize1_1700");
                    }}  
                    disabled='{prize1_1700_save_flag}' 
                    type="button" class="btn btn-info">Generate</button>
                <button
                    on:click={() => {
                        handleSaveGenerator("prize1_1700",prize1_1700);
                    }} 
                    disabled='{prize1_1700_save_flag}' 
                    type="button" class="btn btn-warning">Save</button>
            </div>
		</div>
        <div class="mb-3">
            <label for="exampleForm" class="form-label">
                <div class="d-flex flex-row">
                    <div class="">Prize 20:00</div>
                    <div class="ps-2 mt-3" style="font-size: 11px;color:blue;">OPEN: 19:55 - 20:00</div>
                </div>
            </label>
            <div class="input-group mb-3">
                <Input
                    bind:value={prize1_2000}
                    on:keyup={handleKeyboard_format}    
                    disabled='{prize1_2000_flag}'
                    type="text"
                    minlength="6"
                    maxlength="6"
                    placeholder="Prize 20:00"/>
                <button
                    on:click={() => {
                        generate("prize1_2000");
                    }}  
                    disabled='{prize1_2000_save_flag}' 
                    type="button" class="btn btn-info">Generate</button>
                <button
                    on:click={() => {
                        handleSaveGenerator("prize1_2000",prize1_2000);
                    }} 
                    disabled='{prize1_2000_save_flag}' 
                    type="button" class="btn btn-warning">Save</button>
            </div>
		</div>
        <div class="mb-3">
            <label for="exampleForm" class="form-label">
                <div class="d-flex flex-row">
                    <div class="">Prize 22:00</div>
                    <div class="ps-2 mt-3" style="font-size: 11px;color:blue;">OPEN: 20:55 - 22:00</div>
                </div>
            </label>
            <div class="input-group mb-3">
                <Input
                    bind:value={prize1_2200}
                    on:keyup={handleKeyboard_format}    
                    disabled='{prize1_2200_flag}'
                    type="text"
                    minlength="6"
                    maxlength="6"
                    placeholder="Prize 22:00"/>
                <button
                    on:click={() => {
                        generate("prize1_2200");
                    }}  
                    disabled='{prize1_2200_save_flag}' 
                    type="button" class="btn btn-info">Generate</button>
                <button
                    on:click={() => {
                        handleSaveGenerator("prize1_2200",prize1_2200);
                    }} 
                    disabled='{prize1_2200_save_flag}' 
                    type="button" class="btn btn-warning">Save</button>
            </div>
		</div>
	</slot:template>
</Modal>