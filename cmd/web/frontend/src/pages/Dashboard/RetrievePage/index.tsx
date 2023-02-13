import { Card, Image, Button } from "react-bootstrap";
import { PageHeader } from "./../../../components/layouts/styles";
import { SectionBox } from "./styles";
import React, { useState } from "react";
import {
  createColumnHelper,
  flexRender,
  getCoreRowModel,
  useReactTable,
} from "@tanstack/react-table";
import { Table as BTable } from "react-bootstrap";

import dsRefreshImg from "../../../assets/images/ds-refresh.png";

type Filemetadata = {
  uuid: string;
  fileName: string;
  fileSize: string;
  cid: string;
  actions: string;
};

const defaultData: Filemetadata[] = [
  {
    uuid: "7ec65c5f-2b44-44de-a2f8-890dd6562703",
    fileName: "ProviderInput1.csv",
    fileSize: "1 MB",
    cid: "bafkreicqhqatici5tk3wu4lwdfxj6nmg2pplpu23rxj65x6lebpsdermei",
    actions: "",
  },
  {
    uuid: "006632ec-5677-4343-b7be-9d4ed975f2df",
    fileName: "ProviderInput2.csv",
    fileSize: "10 MB",
    cid: "bafkreicqhqatici5tk3wu4lwdfxj6nmg2pplpu23rxj65x6lebpsdermei",
    actions: "",
  },
];

const columnHelper = createColumnHelper<Filemetadata>();

const columns = [
  columnHelper.accessor("uuid", {
    cell: (info) => info.getValue(),
    header: () => <span>UUID</span>,
  }),
  columnHelper.accessor("fileName", {
    cell: (info) => info.getValue(),
    header: () => <span>File Name</span>,
  }),
  columnHelper.accessor("fileSize", {
    cell: (info) => info.getValue(),
    header: () => <span>File Size</span>,
  }),
  columnHelper.accessor("cid", {
    cell: (info) => info.getValue(),
    header: () => <span>CID</span>,
  }),
  columnHelper.accessor("actions", {
    header: () => "Actions",
    cell: (info) => (
      <>
        <Button style={{ marginRight: 5 }}>Retrieve</Button>
        <Button>Share</Button>
      </>
    ),
  }),
];

const RetrievePage = () => {
  const [data, setData] = React.useState(() => [...defaultData]);

  const table = useReactTable({
    data,
    columns,
    getCoreRowModel: getCoreRowModel(),
  });

  return (
    <>
      <PageHeader>
        <h2>
          <Image className="titleIcon" src={dsRefreshImg} />
          <span>Retrieve</span>
        </h2>
      </PageHeader>
      <SectionBox>
        <Card>
          <Card.Body>
            <BTable bordered hover responsive>
              <thead>
                {table.getHeaderGroups().map((headerGroup) => (
                  <tr key={headerGroup.id}>
                    {headerGroup.headers.map((header) => (
                      <th key={header.id}>
                        {header.isPlaceholder
                          ? null
                          : flexRender(
                              header.column.columnDef.header,
                              header.getContext()
                            )}
                      </th>
                    ))}
                  </tr>
                ))}
              </thead>
              <tbody>
                {table.getRowModel().rows.map((row) => (
                  <tr key={row.id}>
                    {row.getVisibleCells().map((cell) => (
                      <td key={cell.id}>
                        {flexRender(
                          cell.column.columnDef.cell,
                          cell.getContext()
                        )}
                      </td>
                    ))}
                  </tr>
                ))}
              </tbody>
            </BTable>
          </Card.Body>
        </Card>
      </SectionBox>
    </>
  );
};

export default RetrievePage;
