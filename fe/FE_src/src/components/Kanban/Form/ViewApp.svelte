<script>
  import axios from "axios";
  import { Form, FormGroup, Input, Label, Col, Row } from "sveltestrap";

   let rnumber = 0;
   let app_description = "";
   let start_date = "" ;
   let end_date = "";
   let app_permitCreate = "";
   let app_permitOpen = "";
   let app_permitTodo = "";
   let app_permitDoing = "";
   let app_permitDone = "";
   export let app_acronym = "";

   let applications = [];

   async function GetAppData() {
    try {
     const response = await axios.get(`http://localhost:4000/get-application?AppAcronym=${app_acronym}`, { withCredentials: true });
     if (response.data.error) {
      console.log(response.data.error)
     } else if (response) {
      applications = response.data.applications;
      app_acronym = applications.app_acronym;
      app_description = applications.app_description;
      start_date = applications.start_date;
      end_date = applications.end_date;
      rnumber = applications.app_Rnum;
      app_permitCreate = applications.app_permitCreate;
      app_permitOpen = applications.app_permitOpen;
      app_permitTodo = applications.app_permitTodo;
      app_permitDoing = applications.app_permitDoing;
      app_permitDone = applications.app_permitDone;
     }    
    } catch (e) {
      console.log(e)
    }
  }

 $: GetAppData()
</script>

<Form> 
  <Row>
    <Col>
      <FormGroup>
        <Label>App Acronym:</Label>
        <Input type="text" value={app_acronym} readonly/>
      </FormGroup>
    </Col>

    <Col>
      <FormGroup>
        <Label>Running Number:</Label>
        <Input type="number" value={rnumber} readonly/>
      </FormGroup>
    </Col>
  </Row>

    <Row>
      <Col>
        <FormGroup>
          <Label>Description:</Label>
          <Input type="textarea" rows={5} value={app_description} readonly />
        </FormGroup>
      </Col>
    </Row>

    <Row>
      <Col>
        <FormGroup>
          <Label>Create:</Label>
          <Input type= "text" value={app_permitCreate} readonly />
        </FormGroup>
      </Col>

      <Col>
        <FormGroup>
          <Label>Open:</Label>
          <Input type="text" value={app_permitOpen} readonly />
        </FormGroup>
      </Col>

      <Col>
        <FormGroup>
          <Label>To Do:</Label>
          <Input type="text" value={app_permitTodo} readonly />
        </FormGroup>
      </Col>

      <Col>
        <FormGroup>
          <Label>Doing:</Label>
          <Input type="text" value={app_permitDoing} readonly />
        </FormGroup>
      </Col>
      
      <Col>
        <FormGroup>
          <Label>Done:</Label>
          <Input type="text" value={app_permitDone} readonly />
        </FormGroup>
      </Col>
    </Row>

    <Row class="justify-content-md-center">
      <Col xs lg="2">
        <FormGroup>
          <Label>Start:</Label>
          <Input type="date" value={start_date} readonly/>
        </FormGroup>
      </Col>

      <Col xs lg="2">
        <FormGroup>
          <Label>End:</Label>
          <Input type="date" value={end_date} readonly/>
        </FormGroup>
      </Col>
    </Row>
</Form>