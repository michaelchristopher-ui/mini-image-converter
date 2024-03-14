package imageservice

func (i ImageService) RemoveFile(fileName string) {
	i.os.Remove(fileName)
}
