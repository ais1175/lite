import type { ColumnDef, Row, RowSelectionState } from "@tanstack/react-table";
import { useMemo, useState } from "react";
//import { DataTableSkeleton } from "@/components/data-table/data-table-skeleton";
import { DataTable } from "@/components/data/DataTable";
import { Asset } from "@/typings/asset";
import { useParams, useSearchParams } from "react-router";
import { Params } from "@/typings/router";
import { assetColumns } from "./assst-columns";
import { useStorageFiles } from "../api/useStorageFiles";
//import { useToast } from "@fivemanage/ui/hooks/use-toast";
//import { AssetGridSkeleton } from "./asset-grid-skeleton";
//import { AssetGridView } from "./asset-grid-view";

export function AssetList() {
  const [rowSelection, setRowSelection] = useState<RowSelectionState>({});
  const [showDeleteModal, setShowDeleteModal] = useState<boolean>(false);

  const columns = useMemo<ColumnDef<Asset>[]>(() => assetColumns(), []);

  //const { toast } = useToast();
  //const trpcUtils = api.useUtils();

  const [searchParams] = useSearchParams();
  const params = useParams<Params>();

  const view = searchParams?.get("view") ?? "list";
  //const assetType = searchParams?.get("type") ?? "all";
  //const fromDate = searchParams.get("from") ?? new Date().toString();
  //const toDate = searchParams.get("to") ?? new Date().toString();
  //const page = searchParams.get("page");
  const search = searchParams.get("search");

  /*const { data, isLoading } = api.assets.list.useQuery({
    //toDate,
    //fromDate,
    search: search,
    organizationId: params.organizationId as string,
    //currentPage: Number(page),
    //type: assetType !== "all" ? assetType : undefined,
  }); */

  const { data } = useStorageFiles({
    organizationId: params.organizationId as string,
    search: search ?? undefined,
  });

  /*const { mutateAsync, isPending: isDeleteLoading } =
    api.assets.deleteMany.useMutation({
      onSuccess() {
        setShowDeleteModal(false);
        setRowSelection({});

        //void trpcUtils.assets.list.invalidate();
        //router.refresh();

        toast({
          title: "Success",
          description: "Assets successfully deleted.",
        }); 
      },
      onError: () => {
        setShowDeleteModal(false);
        setRowSelection({});

        void trpcUtils.assets.list.invalidate();
        router.refresh();

        toast({
          title: "Error",
          description: "Failed to delete assets. Please try again later.",
        }); 
      },
    }); */

  /*const handleDelete = async (rows: Row<Asset>[]) => {
    try {
      await mutateAsync({
        teamId: queryParams.teamId as string,
        files: rows.map((r) => {
          return {
            id: r.original.id,
            key: r.original.key ?? "",
            type: r.original.type,
          };
        }),
      });
    } catch (err) {
      console.error(err);
    }
  }; */

  /*if (isLoading) {
    return view === "list" ? (
      <DataTableSkeleton
        columnCount={columns.length - 1}
        rowCount={10}
        showViewOptions={false}
      />
    ) : (
      {
        <AssetGridSkeleton rowCount={6} />
      }
    );
  } */

  return (
    <div>
      <div>
        {/*data && view === "grid" && (
          <AssetGridView
            assets={data.uploads}
            currentPage={page}
            count={data.totalCount}
          />
        )*/}
      </div>
      <div>
        {data && view === "list" && (
          <DataTable
            columns={columns}
            rowSelection={rowSelection}
            setRowSelection={setRowSelection}
            showDeleteModal={showDeleteModal}
            setShowDeleteModal={setShowDeleteModal}
            isLoading={false}
            data={data}
            totalCount={data.length}
            onDelete={() => {}}
          />
        )}
      </div>
    </div>
  );
}
