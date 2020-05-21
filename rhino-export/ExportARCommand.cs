using System;
using System.IO;
using Rhino;
using Rhino.Commands;
using Rhino.Geometry;
using Rhino.FileIO;
using Rhino.Input.Custom;
using System.Net.Http;
using System.Threading.Tasks;
using Rhino.UI;

namespace ExportAR
{
    public class ExportARCommand : Command
    {
        public ExportARCommand()
        {
            // Rhino only creates one instance of each command class defined in a
            // plug-in, so it is safe to store a refence in a static property.
            Instance = this;
        }

        ///<summary>The only instance of this command.</summary>
        public static ExportARCommand Instance
        {
            get; private set;
        }

        ///<returns>The command name as it appears on the Rhino command line.</returns>
        public override string EnglishName
        {
            get { return "ExportAR"; }
        }

        protected override Result RunCommand(RhinoDoc doc, RunMode mode)
        {
            GetObject go = new GetObject();
            go.SetCommandPrompt("Please select the geometry you want to export to AR");
            go.GetMultiple(0, 0);
            
            if (go.ObjectCount > 0)
            {
                string docName = doc.Name.Replace(".3dm", "");

                //Export
                string dir = Path.Combine(Environment.GetFolderPath(Environment.SpecialFolder.MyDocuments), "AR",
                    "Export");
                string objPath = Path.Combine(Environment.GetFolderPath(Environment.SpecialFolder.MyDocuments), "AR",
                    "Export", docName + ".obj");
                string mtlPath = Path.Combine(Environment.GetFolderPath(Environment.SpecialFolder.MyDocuments), "AR",
                    "Export", docName + ".mtl");
                string encodedPath = Path.Combine(Environment.GetFolderPath(Environment.SpecialFolder.MyDocuments), "AR",
                     "Export", docName + ".txt");

                Directory.CreateDirectory(dir);
                FileUtilties.DeleteFiles(dir);
                FileObj.Write(objPath, doc, CreateObjWriteOptions());
                //FileUtilties.ModifyMtl(mtlPath);

                byte[] objFile = File.ReadAllBytes(objPath);
                byte[] mtlFile = File.ReadAllBytes(mtlPath);


                UploadToServer(objFile, mtlFile, docName);

                doc.Objects.UnselectAll();

            }

            return Result.Success;
        }

        /// <summary>
        /// Generate standard obj file write options
        /// </summary>
        /// <returns>A FileObjWriteOptions object containing the settings for the export.</returns>
        private static FileObjWriteOptions CreateObjWriteOptions()
        {
            //Create general export settings
            FileWriteOptions writeOptions = new FileWriteOptions
            {
                WriteSelectedObjectsOnly = true,
                WriteGeometryOnly = true
            };

            //Create obj export settings for everything other than rooms
            FileObjWriteOptions objWriteOptions = new FileObjWriteOptions(writeOptions)
            {
                CreateNgons = false,
                ExportMaterialDefinitions = true,
                ExportNormals = true,
                ExportTcs = true,
                ExportVcs = true,
                SignificantDigits = 4,
                ExportGroupNameLayerNames = FileObjWriteOptions.ObjGroupNames.NoGroups,
                UseRelativeIndexing = false,
                MeshParameters = MeshingParameters.FastRenderMesh,
                MapZtoY = true,
                UnderbarMaterialNames = true
            };

            return objWriteOptions;
        }

        private async Task<string> UploadToServer(byte[] objFile, byte[] mtlFile, String docName)
        {
            HttpClient httpClient = new HttpClient();
            MultipartFormDataContent form = new MultipartFormDataContent();
            form.Add(new ByteArrayContent(objFile, 0, objFile.Length), "doc.obj", docName + ".obj");
            form.Add(new ByteArrayContent(mtlFile, 0, mtlFile.Length), "doc.mtl", docName + ".mtl");

            string url = "https://ar.portfo.io/?mode=obj";

            HttpResponseMessage response = await httpClient.PostAsync(url, form);

            response.EnsureSuccessStatusCode();
            httpClient.Dispose();
            string sd = response.Content.ReadAsStringAsync().Result;
            string frontendUrl = "https://ar-viewer.netlify.app/";
            Dialogs.ShowMessage($"Please search for {sd} on {frontendUrl}", "Upload Success!");
            RhinoApp.WriteLine($"Please search for {sd} on {frontendUrl}", "Upload Success!");
            return sd;
        }
    }
}
