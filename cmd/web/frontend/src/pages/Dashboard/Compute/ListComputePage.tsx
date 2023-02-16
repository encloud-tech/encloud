import { Card } from "react-bootstrap";
import { PageHeader } from "./../../../components/layouts/styles";
import {
  createColumnHelper,
  flexRender,
  getCoreRowModel,
  useReactTable,
} from "@tanstack/react-table";
import { Table as BTable } from "react-bootstrap";
import { useState } from "react";
import { SectionBox } from "./styles";
import { Link } from "react-router-dom";

type Computedata = {
  id: string;
  jobId: string;
  state: string;
  created: string;
  verified: string;
  published: string;
};

const defaultData: Computedata[] = [
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

const columnHelper = createColumnHelper<Computedata>();

const columns = [
  columnHelper.accessor("id", {
    cell: (info) => info.getValue(),
    header: () => <span>ID</span>,
  }),
  columnHelper.accessor("jobId", {
    cell: (info) => info.getValue(),
    header: () => <span>Job ID</span>,
  }),
  columnHelper.accessor("state", {
    cell: (info) => info.getValue(),
    header: () => <span>State</span>,
  }),
  columnHelper.accessor("created", {
    cell: (info) => info.getValue(),
    header: () => <span>Created</span>,
  }),
  columnHelper.accessor("verified", {
    cell: (info) => info.getValue(),
    header: () => <span>Verified</span>,
  }),
  columnHelper.accessor("published", {
    cell: (info) => info.getValue(),
    header: () => <span>Published</span>,
  }),
  columnHelper.accessor("id", {
    cell: (info) => (
      <Link to={`/get-results/${info.getValue()}`} className="btn btn-primary">
        Get Results
      </Link>
    ),
    header: () => <span>Actions</span>,
  }),
];

const ListComputePage = () => {
  const [data, setData] = useState(() => [...defaultData]);

  const table = useReactTable({
    data,
    columns,
    getCoreRowModel: getCoreRowModel(),
  });

  return (
    <>
      <PageHeader>
        <h2>
          <span>List</span>
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

export default ListComputePage;
