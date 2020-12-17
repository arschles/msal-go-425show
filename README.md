# MSAL, Cognitive Services and Go

Examples of using Azure Cognitive Services, MSAL (Microsoft Authentication Library, Formerly Azure Active Directory), and Go.

We'll be building a CLI client for the [425 show](https://www.twitch.tv/425show) to take images of receipts, etc..., send them to [Azure Cognitive Service](https://azure.microsoft.com/en-us/services/cognitive-services/)'s [OCR APIs](https://docs.microsoft.com/en-us/azure/cognitive-services/computer-vision/concept-recognizing-text), and get text summaries of a receipt's information like the vendor, amount paid, and so on.

We'll start off with [this example code](https://github.com/Azure-Samples/cognitive-services-quickstart-code/blob/master/go/ComputerVision/ComputerVisionQuickstart.go) and:

- Explain what's going on
- Refactor out the calls to other APIs besides OCR
- Write the code to read in the receipt images and integrate with the OCR call
- Write code to use MSAL rather than the [subscription key](https://github.com/Azure-Samples/cognitive-services-quickstart-code/blob/ee64dd51ebee99a34db12748e15ed23a536e47e2/go/ComputerVision/ComputerVisionQuickstart.go#L104) approach to authenticate with the API
  - We can use [this library](https://github.com/AzureAD/microsoft-authentication-library-for-go) for MSAL with Go


Additional resources:

- [MSAL conceptual overview](https://docs.microsoft.com/en-us/azure/active-directory/develop/msal-overview)
- [Step by step getting started guide](https://docs.microsoft.com/en-us/azure/cognitive-services/authentication?tabs=powershell#authenticate-with-azure-active-directory) - this is in PowerShell so a lot of the work involved will be translating this to raw HTTP calls, or MSAL SDK calls
