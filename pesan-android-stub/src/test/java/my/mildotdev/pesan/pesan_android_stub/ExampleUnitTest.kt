package my.mildotdev.pesan.pesan_android_stub

import com.github.miltsm.pesan_grpc_stubs.PublicGrpcKt
import com.github.miltsm.pesan_grpc_stubs.passwordLoginRequest
import com.google.protobuf.ByteString
import io.grpc.ManagedChannelBuilder
import kotlinx.coroutines.Dispatchers
import kotlinx.coroutines.asExecutor
import kotlinx.coroutines.runBlocking
import org.junit.Assert.assertEquals
import org.junit.Test
import java.net.URL

/**
 * Example local unit test, which will execute on the development machine (host).
 *
 * See [testing documentation](http://d.android.com/tools/testing).
 */
class ExampleUnitTest {
    @Test
    fun addition_isCorrect() {
        assertEquals(4, 2 + 2)
    }

    @Test
    fun test_grcp_connection() {
        val uri = URL("http://{replace-your-ip}:50051")
        val channel = let {
            val builder = ManagedChannelBuilder.forAddress(uri.host, uri.port)
//            if (uri. == "https")
//                builder.useTransportSecurity()
//            else
            builder.usePlaintext()

            builder.executor(Dispatchers.IO.asExecutor()).build()
        }

        val pesan = PublicGrpcKt.PublicCoroutineStub(channel)
        val request = passwordLoginRequest {
            this.userHandle = "tested"
            this.password = ByteString.copyFromUtf8("testtesttesT1#")
        }
        try {
            runBlocking {
                val result = pesan.loginWithPassword(request)
                assertEquals(request.userHandle, result.userHandle)
            }
        } catch (e: Exception) {
            assert(false)
        }
    }
}