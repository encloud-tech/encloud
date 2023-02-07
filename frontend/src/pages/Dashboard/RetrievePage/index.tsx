import { Card, Image, Button } from "react-bootstrap";
import { PageHeader } from "./../../../components/layouts/styles";
import { SectionBox, TableCell } from "./styles";
import React, { useState } from "react";
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

import dsRefreshImg from "../../../assets/images/ds-refresh.png";

const dataSets = [
  {
    uuid: "7ec65c5f-2b44-44de-a2f8-890dd6562703",
    fileName: "ProviderInput1.csv",
    fileSize: "1 MB",
    cid: "bafkreicqhqatici5tk3wu4lwdfxj6nmg2pplpu23rxj65x6lebpsdermei",
    actions: "abc",
  },
  {
    uuid: "006632ec-5677-4343-b7be-9d4ed975f2df",
    fileName: "ProviderInput2.csv",
    fileSize: "10 MB",
    cid: "bafkreicqhqatici5tk3wu4lwdfxj6nmg2pplpu23rxj65x6lebpsdermei",
    actions: "",
  },
];

const RetrievePage = () => {
  const [columns] = useState([
    { name: "uuid", title: "UUID" },
    { name: "fileName", title: "File Name" },
    { name: "fileSize", title: "File Size" },
    { name: "cid", title: "CID" },
    { name: "actions", title: "Actions" },
  ]);
  const [rows] = useState(dataSets);

  const Cell = (props: any) => {
    const { column } = props;
    if (column.name === "actions") {
      return (
        <Table.Cell {...props}>
          <Button style={{ marginRight: 5 }}>Retrieve</Button>
          <Button>Share</Button>
        </Table.Cell>
      );
    }
    return <TableCell className="tdbox" {...props} />;
  };

  return (
    <>
      <PageHeader>
        <h2>
          <Image className="titleIcon" src={dsRefreshImg} />
          <span>Retrieve</span>
        </h2>
      </PageHeader>
      <SectionBox>
        <Card className="mb-3">
          <Card.Body className="p-0">
            <Grid rows={rows} columns={columns}>
              <PagingState defaultCurrentPage={0} pageSize={10} />
              <IntegratedPaging />
              <SortingState
                defaultSorting={[{ columnName: "uuid", direction: "asc" }]}
              />
              <IntegratedSorting />
              <Table cellComponent={Cell} />
              <TableHeaderRow showSortingControls />
              <PagingPanel />
            </Grid>
          </Card.Body>
        </Card>
      </SectionBox>
    </>
  );
};

export default RetrievePage;
