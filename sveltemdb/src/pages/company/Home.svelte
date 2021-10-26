<script>
    import Panel from "../../components/Panel.svelte";
	import Modal from "../../components/Modal.svelte";

	export let table_header_font
	export let table_body_font
	export let totalrecord = 0
	let title_page = "Home"
</script>
<div class="container" style="margin-top: 70px;">
    <div class="row">
        <div class="col-sm-12">
            <button
                type="button"
                class="btn btn-dark btn-sm"
                data-mdb-toggle="modal"
                data-mdb-target="#exampleModal">
                NEW
            </button>
            <button
                type="button"
                class="btn btn-primary btn-sm">
                Refresh
            </button>
            <Panel
                card_title="{title_page}"
                card_footer={totalrecord}>
                <slot:template slot="card-body">
                    <table class="table table-striped table-sm table-light border-dark">
                        <thead>
                            <tr>
                                <th width="1%" style="text-align: center;vertical-align: top;">&nbsp;</th>
                                <th width="1%" style="text-align: center;vertical-align: top;font-weight:bold;font-size:{table_header_font};">NO</th>
                                <th width="*" style="text-align: left;vertical-align: top;font-weight:bold;font-size: {table_header_font};">NAME</th>
                                <th width="20%" style="text-align: left;vertical-align: top;font-weight:bold;font-size: {table_header_font};">EMAIL</th>
                            </tr>
                        </thead>
                        <tbody>
                            <tr>
                                <td style="text-align: center;vertical-align: top;cursor:pointer;"><i class="far fa-edit"></i></td>
                                <td style="text-align: center;vertical-align: top;font-size: {table_body_font};">1</td>
                                <td style="text-align: left;vertical-align: top;font-size: {table_body_font};">BUDI</td>
                                <td style="text-align: left;vertical-align: top;font-size: {table_body_font};">BUDI@gmail.com</td>
                            </tr>
                            <tr>
                                <td style="text-align: center;vertical-align: top;cursor:pointer;"><i class="far fa-edit"></i></td>
                                <td style="text-align: center;vertical-align: top;font-size: {table_body_font};">1</td>
                                <td style="text-align: left;vertical-align: top;font-size: {table_body_font};">BUDI</td>
                                <td style="text-align: left;vertical-align: top;font-size: {table_body_font};">BUDI@gmail.com</td>
                            </tr>
                            <tr>
                                <td style="text-align: center;vertical-align: top;cursor:pointer;"><i class="far fa-edit"></i></td>
                                <td style="text-align: center;vertical-align: top;font-size: {table_body_font};">1</td>
                                <td style="text-align: left;vertical-align: top;font-size: {table_body_font};">BUDI</td>
                                <td style="text-align: left;vertical-align: top;font-size: {table_body_font};">BUDI@gmail.com</td>
                            </tr>
                            <tr>
                                <td style="text-align: center;vertical-align: top;cursor:pointer;"><i class="far fa-edit"></i></td>
                                <td style="text-align: center;vertical-align: top;font-size: {table_body_font};">1</td>
                                <td style="text-align: left;vertical-align: top;font-size: {table_body_font};">BUDI</td>
                                <td style="text-align: left;vertical-align: top;font-size: {table_body_font};">BUDI@gmail.com</td>
                            </tr>
                            <tr>
                                <td style="text-align: center;vertical-align: top;cursor:pointer;"><i class="far fa-edit"></i></td>
                                <td style="text-align: center;vertical-align: top;font-size: {table_body_font};">1</td>
                                <td style="text-align: left;vertical-align: top;font-size: {table_body_font};">BUDI</td>
                                <td style="text-align: left;vertical-align: top;font-size: {table_body_font};">BUDI@gmail.com</td>
                            </tr>
                            <tr>
                                <td style="text-align: center;vertical-align: top;cursor:pointer;"><i class="far fa-edit"></i></td>
                                <td style="text-align: center;vertical-align: top;font-size: {table_body_font};">1</td>
                                <td style="text-align: left;vertical-align: top;font-size: {table_body_font};">BUDI</td>
                                <td style="text-align: left;vertical-align: top;font-size: {table_body_font};">BUDI@gmail.com</td>
                            </tr>
                            <tr>
                                <td style="text-align: center;vertical-align: top;cursor:pointer;"><i class="far fa-edit"></i></td>
                                <td style="text-align: center;vertical-align: top;font-size: {table_body_font};">1</td>
                                <td style="text-align: left;vertical-align: top;font-size: {table_body_font};">BUDI</td>
                                <td style="text-align: left;vertical-align: top;font-size: {table_body_font};">BUDI@gmail.com</td>
                            </tr>
                            <tr>
                                <td style="text-align: center;vertical-align: top;cursor:pointer;"><i class="far fa-edit"></i></td>
                                <td style="text-align: center;vertical-align: top;font-size: {table_body_font};">1</td>
                                <td style="text-align: left;vertical-align: top;font-size: {table_body_font};">BUDI</td>
                                <td style="text-align: left;vertical-align: top;font-size: {table_body_font};">BUDI@gmail.com</td>
                            </tr>
                        </tbody>
                    </table>
                </slot:template>
            </Panel>
        </div>
    </div>
</div>

<Modal
	modal_id="exampleModal"
	modal_size="modal-dialog-centered"
	modal_title="Title"
	modal_footer={true}>
	<slot:template slot="body">
		<div class="mb-3">
			<input type="type" class="form-control" id="username" placeholder="Username">
		</div>
		<div class="mb-3">
			<input type="password" class="form-control" id="password" placeholder="Password">
		</div>
	</slot:template>
	<slot:template slot="footer">
		<button type="button" class="btn btn-dark">Save</button>
	</slot:template>
</Modal>