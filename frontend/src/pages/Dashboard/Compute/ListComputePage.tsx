import { Card } from "react-bootstrap";
import { PageHeader } from "./../../../components/layouts/styles";
import {
  SortingState,
  IntegratedSorting,
  PagingState,
  IntegratedPaging,
} from "@devexpress/dx-react-grid";
import {
  Grid,
  Table,
  TableHeaderRow,
  PagingPanel,
} from "@devexpress/dx-react-grid-bootstrap4";
import "@devexpress/dx-react-grid-bootstrap4/dist/dx-react-grid-bootstrap4.css";
import { useState } from "react";

const dataSets = [
  {
    id: "5e488f74",
    job: "Docker ubuntu echo T..",
    state: "Completed",
    created: "22:10:05",
    verified: "-",
    published: "/ipfs/QmX6fQ9LXfaKve...",
  },
];

const ListComputePage = () => {
  const [columns] = useState([
    { name: "id", title: "ID" },
    { name: "job", title: "Job" },
    { name: "state", title: "State" },
    { name: "created", title: "Created" },
    { name: "verified", title: "Verified" },
    { name: "published", title: "Published" },
  ]);
  const [rows] = useState(dataSets);

  return (
    <>
      <PageHeader>
        <h2>
          <span>List</span>
        </h2>
      </PageHeader>

      <Card>
        <Card.Body>
          <Grid rows={rows} columns={columns}>
            <PagingState defaultCurrentPage={0} pageSize={10} />
            <IntegratedPaging />
            <SortingState
              defaultSorting={[{ columnName: "uuid", direction: "asc" }]}
            />
            <IntegratedSorting />
            <Table />
            <TableHeaderRow showSortingControls />
            <PagingPanel />
          </Grid>
        </Card.Body>
      </Card>
    </>
  );
};

export default ListComputePage;
