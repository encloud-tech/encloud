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
import { TableCell } from "../RetrievePage/styles";
import { Link } from "react-router-dom";

const dataSets = [
  {
    id: "47805f5c",
    jobId: "e3f8c209-d683-4a41-b840-f09b88d087b9",
    state: "Completed",
    created: "12:11:09",
    verified: "-",
    published: "/ipfs://bafybeig7mdkzcgpacp...",
  },
  {
    id: "ebd9bf2f",
    jobId: "51225160-807e-48b8-88c9-28311c7899e1",
    state: "Pending",
    created: "22:10:05",
    verified: "-",
    published: "/ipfs/QmX6fQ9LXfaKve...",
  },
];

const ListComputePage = () => {
  const [columns] = useState([
    { name: "id", title: "ID" },
    { name: "jobId", title: "Job ID" },
    { name: "state", title: "State" },
    { name: "created", title: "Created" },
    { name: "verified", title: "Verified" },
    { name: "published", title: "Published" },
    { name: "actions", title: "Actions" },
  ]);

  const [rows] = useState(dataSets);

  const Cell = (props: any) => {
    const { column, row } = props;
    if (column.name === "actions") {
      return (
        <Table.Cell {...props}>
          <Link
            to={`/dashboard/get-results/${row.jobId}`}
            className="btn btn-primary"
          >
            Get Results
          </Link>
        </Table.Cell>
      );
    }
    return <TableCell className="tdbox" {...props} />;
  };

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
              defaultSorting={[{ columnName: "id", direction: "asc" }]}
            />
            <IntegratedSorting />
            <Table cellComponent={Cell} />
            <TableHeaderRow showSortingControls />
            <PagingPanel />
          </Grid>
        </Card.Body>
      </Card>
    </>
  );
};

export default ListComputePage;
