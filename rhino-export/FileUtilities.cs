using System.IO;

namespace ExportAR
{
    public class FileUtilties

    {

        /// <summary>
        /// Deletes all files in the directory.
        /// </summary>
        /// <param name="dir">The directory to delete files in.</param>
        public static void DeleteFiles(string dir)
        {
            foreach (string file in Directory.EnumerateFiles(dir)) File.Delete(file);
        }

    }
}